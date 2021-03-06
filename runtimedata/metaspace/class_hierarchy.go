package metaspace

// 判断类之间的关系

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
func (c *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, c
	if s == t {
		return true
	}
	if !s.IsArray() {  //不是数组
		if !s.IsInterface() { // 不是接口
			if !t.IsInterface() {
				return s.IsSubClassOf(t)
			} else {
				s.IsImplements(t) // 判断s是否实现t
			}

		}else {
			//s是接口
			if !t.IsInterface(){
				// t不是接口
				return t.isJlObject()  //是否是Object
			} else{
				return t.isJlCloneable() || t.isJlSerializable()
			}
		}
	}else {
		// s是数组
		if !t.IsArray() {
			// t不是数组
			if !t.IsInterface() {
				return t.isJlObject()
			}else {
				return t.isJlCloneable() || t.isJlSerializable()
			}
		}else{
			// t是数组
			// 等待完成
			return false
		}
	}
	return false
}

// 判断c是否继承于某个类
// 递归往上找 找到一个c的一个父类等于other
func (c *Class) IsSubClassOf(other *Class) bool {
	for class := c.superClass; class != nil; class = class.superClass {
		if class == other {
			return true
		}
	}
	return false
}

// 判断当前类是否实现一个接口
func (c *Class) IsImplements(iface *Class) bool {
	for class := c; class != nil; class = c.superClass { // 循环遍历当前类和父类
		for _, i := range class.interfaces { // 循环遍历接口
			if i == iface || i.isSubInterfaceOf(iface) { // 进行查找
				return true
			}
		}
	}
	return false
}

// 判断当前是否继承自接口
func (c *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range c.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) { // 先查找自己的接口 然后递归查询接口继承的接口
			return true // 接口允许继承多个接口
		}
	}
	return false
}

// 判断c是不是other的超类
func (c *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(c)
}