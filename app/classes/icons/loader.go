package icons

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/starter-go/application"
	"github.com/starter-go/application/resources"
	"github.com/starter-go/mimetypes"
	"github.com/starter-go/vlog"
)

// Loader ...
type Loader struct {

	//starter:component

	_as func(mimetypes.Registry, Service) //starter:as(".","#")

	ConfigProvider ConfigProvider      //starter:inject("#")
	AppContext     application.Context //starter:inject("context")

	cache *cache
}

func (inst *Loader) _impl() (mimetypes.Registry, Service) {
	return inst, inst
}

// ListRegistrations ...
func (inst *Loader) ListRegistrations() []*mimetypes.Registration {
	c := inst.getCache()
	return c.listRegistrations()
}

func (inst *Loader) load1(c *cache) {
	err := inst.load2(c)
	if err == nil {
		return
	}
	vlog.Warn(err.Error())
}

func (inst *Loader) load2(c *cache) error {

	resset := inst.AppContext.GetResources()
	all := resset.Paths()
	cfg := inst.ConfigProvider.Configuration()
	prefix := cfg.ResPathPrefix

	c.config = cfg
	c.init()

	for _, path := range all {
		if strings.HasPrefix(path, prefix) {
			res, err := resset.GetResource(path)
			if err != nil {
				vlog.Warn(err.Error())
				continue
			}
			if res == nil {
				continue
			}
			err = c.addRes(res)
			if err != nil {
				vlog.Warn(err.Error())
				continue
			}
		}
	}
	return nil
}

func (inst *Loader) getCache() *cache {
	c := inst.cache
	if c == nil {
		c = &cache{}
		inst.load1(c)
		inst.cache = c
	}
	return c
}

// FindImage ...
func (inst *Loader) FindImage(c context.Context, webpath string) (*Image, error) {
	ca := inst.getCache()
	mt := ca.webpath2type(webpath)
	return ca.loadImgae(c, mt)
}

////////////////////////////////////////////////////////////////////////////////

type cache struct {
	config           *Configuration
	mutex            sync.Mutex
	table            map[mimetypes.Type]*imageHolder // map[type] holder (as PK)
	tableNameIndexer map[string]*imageHolder         // map[name] holder (as indexer)
}

func (inst *cache) init() {
	inst.table = make(map[mimetypes.Type]*imageHolder)
}

func (inst *cache) getTableNameIdx() map[string]*imageHolder {
	idx := inst.tableNameIndexer
	if idx != nil {
		return idx
	}
	// load indexer
	src := inst.table
	idx = make(map[string]*imageHolder)
	for _, item := range src {
		name := item.resFileName
		idx[name] = item
	}
	inst.tableNameIndexer = idx
	return idx
}

func (inst *cache) listRegistrations() []*mimetypes.Registration {

	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()

	cfg := inst.config
	src := inst.table
	dst := make([]*mimetypes.Registration, 0)

	for _, h := range src {
		mtype := h.mediaType
		filename := h.res.SimpleName()
		item := &mimetypes.Registration{
			Name:     mtype.String(),
			Priority: cfg.Priority,
		}
		item.Info.Type = mtype
		item.Info.Icon = inst.type2webpath(mtype, filename)
		dst = append(dst, item)
	}
	return dst
}

func (inst *cache) type2webpath(t mimetypes.Type, filename string) string {
	cfg := inst.config
	path := cfg.WebPathPrefix + "/" + t.Pure() + "/" + filename
	elist := path2elements(path)
	elist = normalizePathEl(elist)
	return elements2path(elist)
}

func (inst *cache) webpath2type(path string) mimetypes.Type {
	prefix := inst.config.WebPathPrefix
	if strings.HasPrefix(path, prefix) {
		elist := path2elements(path[len(prefix):])
		elist = normalizePathEl(elist)
		if len(elist) >= 2 {
			str := elist[0] + "/" + elist[1]
			return mimetypes.Type(str)
		}
	}
	return "bad/type"
}

func (inst *cache) parseMediaTypeWithResFileName(filename string) mimetypes.Type {
	i1 := strings.IndexRune(filename, '-')
	i2 := strings.LastIndexByte(filename, '.')
	if 0 < i1 && i1 < i2 {
		str1 := filename[0:i1]
		str2 := filename[i1+1 : i2]
		str := str1 + "/" + str2
		return mimetypes.Type(str)
	}
	return mimetypes.Type(filename) // "bad/type"
}

func (inst *cache) addRes(res resources.Resource) error {
	mt := inst.parseMediaTypeWithResFileName(res.SimpleName())
	h := &imageHolder{
		cache:       inst,
		resPath:     res.Path(),
		resFileName: res.SimpleName(),
		res:         res,
		key:         mt,
		mediaType:   mt,
	}
	inst.table[h.mediaType] = h
	return nil
}

func (inst *cache) loadImgae(_ context.Context, mtype mimetypes.Type) (*Image, error) {

	locker := &inst.mutex
	locker.Lock()
	defer locker.Unlock()

	item := inst.table[mtype]
	if item == nil {
		return nil, fmt.Errorf("no type: " + mtype.String())
	}

	data, err := item.getImageData(5)
	if err != nil {
		return nil, err
	}

	img := &Image{
		Name: mtype.Pure(),
		Path: mtype.Full(),
		Type: item.getImageType(),
		Data: data,
	}
	return img, nil
}

////////////////////////////////////////////////////////////////////////////////

type imageHolder struct {
	key       mimetypes.Type
	mediaType mimetypes.Type

	cache       *cache
	resPath     string
	resFileName string
	res         resources.Resource
	data        []byte
}

func (inst *imageHolder) getImageType() mimetypes.Type {
	const bad = mimetypes.Type("bad/type")
	name := strings.ToLower(inst.resFileName)
	i9 := strings.LastIndexByte(name, '.')
	if i9 < 0 {
		return bad
	}
	suffix := name[i9:]
	switch suffix {
	case ".png":
		return "image/png"
	case ".svg":
		return "image/svg"
	case ".jpg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	}
	return bad
}

func (inst *imageHolder) getImageData(ttl int) ([]byte, error) {
	d := inst.data
	if d != nil {
		return d, nil
	}
	d, err := inst.loadImageData(ttl)
	if err != nil {
		return nil, err
	}
	inst.data = d
	return d, nil
}

// ttl （time to alive） 用于防止无限循环引用
func (inst *imageHolder) loadImageData(ttl int) ([]byte, error) {

	if ttl < 0 {
		return nil, fmt.Errorf("load image with recursion")
	}
	ttl--
	data, err := inst.res.ReadBinary()
	if err != nil {
		return nil, err
	}

	islink, filename := inst.isLink(data)
	if islink {
		return inst.loadAsLink(filename, ttl)
	}
	return data, nil
}

func (inst *imageHolder) loadAsLink(filename string, ttl int) ([]byte, error) {
	indexer := inst.cache.getTableNameIdx()
	next := indexer[filename]
	if next == nil {
		return nil, fmt.Errorf("no icon image file with name: [%s]", filename)
	}
	return next.getImageData(ttl)
}

func (inst *imageHolder) isLink(data []byte) (islink bool, filename string) {
	size := len(data)
	if size > 256 {
		return false, ""
	}

	// check chars
	for _, b := range data {
		if b < '\r' {
			return false, ""
		} else if b > 127 {
			return false, ""
		} else if b == '<' || b == '>' {
			return false, ""
		}
	}

	islink = true
	filename = strings.TrimSpace(string(data))
	return
}

////////////////////////////////////////////////////////////////////////////////
