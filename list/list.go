package list

type Any interface {
	int | string
}

// Remove 删除a切片中的b元素 只删除某一个元素会修改切片的顺序
func Remove[T Any](a, b []T) []T {
	m := make(map[T]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		delete(m, v)
	}
	newSli := make([]T, 0)
	for k, _ := range m {
		newSli = append(newSli, k)
	}
	return newSli
}

// Add 像切片中添加元素不重复
func Add[T Any](a, b []T) []T {
	a = append(a, b...)
	m := make(map[T]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}
	newSli := make([]T, 0)
	for k, _ := range m {
		newSli = append(newSli, k)
	}
	return newSli
}

// Contain 判断切片中是否包含某个、某些元素
func Contain[T Any](a []T, b ...T) bool {
	m := make(map[T]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}
	flag := false
	for _, v := range b {
		if _, flag = m[v]; !flag {
			return flag
		}
	}
	return flag
}

// IsEmpty 判断切片是否为空 如果为空则返回true 否则返回false
func IsEmpty[T Any](a []T) bool {
	return len(a) == 0
}

// Index 返回该元素再数组中第一次出现的位置
func Index[T Any](a []T, b T) int {
	for i := range a {
		if a[i] == b {
			return i
		}
	}
	return -1
}
