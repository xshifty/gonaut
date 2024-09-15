package gonaut

// LazyLoadFunc base type for lazy load function container
type LazyLoadFunc func(c Container) interface{}

// Dependency holds lazy loader function and dependency value
type Dependency struct {
	lazyLoader LazyLoadFunc
	dependency interface{}
}

// Container store dependencies
type Container map[string]Dependency

var c = Container{}

// Set add a new dependency to container
func Set(key string, loader LazyLoadFunc) Container {
	return c.Set(key, loader)
}

// Get a dependency from container
func Get(key string) interface{} {
	return c.Get(key)
}

// Set from container struct will register a dependency
func (c Container) Set(key string, loader LazyLoadFunc) Container {
	if _, ok := c[key]; !ok {
		c[key] = Dependency{
			lazyLoader: loader,
			dependency: nil,
		}
	}

	return c
}

// Get method from container struct will load or get dependency
func (c Container) Get(key string) interface{} {
	dep, ok := c[key]
	if !ok {
		return nil
	}
	if dep.dependency == nil {
		dep.dependency = dep.lazyLoader(c)
	}

	return dep.dependency
}
