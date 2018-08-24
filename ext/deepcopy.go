package ext

// . "reflect"

// func DeepCopy(obj interface{}) interface{} {
// 	return deepcopy(NewValue(obj), map[uintptr]map[Type]Value{}).Interface()
// }

// func deepcopy(obj Value, visited map[uintptr]map[Type]Value) Value {
// 	if obj == nil {
// 		return nil
// 	}
// 	n := MakeZero(obj.Type())

// 	addr := obj.Addr()
// 	if omap := visited[addr]; omap != nil {
// 		old := omap[obj.Type()]
// 		if old != nil {
// 			return old
// 		}
// 		omap[obj.Type()] = obj
// 	} else {
// 		visited[addr] = map[Type]Value{obj.Type(): obj}
// 	}

// 	switch ov := obj.(type) {
// 	case *ArrayValue:
// 		nv := n.(*ArrayValue)
// 		for i := 0; i < ov.Len(); i++ {
// 			nv.Elem(i).SetValue(deepcopy(ov.Elem(i), visited))
// 		}
// 	case *BoolValue, *ChanValue, *ComplexValue, *FloatValue, *FuncValue, *IntValue, *StringValue, *UintValue, *UnsafePointerValue:
// 		n.SetValue(ov)
// 	case *InterfaceValue:
// 		nv := n.(*InterfaceValue)
// 		nv.SetValue(ov)
// 		nv.Elem().SetValue(deepcopy(ov.Elem(), visited))
// 	case *MapValue:
// 		nv := n.(*MapValue)
// 		nv.SetValue(MakeMap(obj.Type().(*MapType)))
// 		for _, k := range ov.Keys() {
// 			nkey := deepcopy(k, visited)
// 			nval := deepcopy(ov.Elem(k), visited)
// 			nv.SetElem(nkey, nval)
// 		}
// 	case *PtrValue:
// 		n.(*PtrValue).PointTo(deepcopy(ov.Elem(), visited))
// 	case *SliceValue:
// 		nv := n.(*SliceValue)
// 		nv.SetValue(MakeSlice(obj.Type().(*SliceType), ov.Len(), ov.Cap()))
// 		for i := 0; i < ov.Len(); i++ {
// 			nv.Elem(i).SetValue(deepcopy(ov.Elem(i), visited))
// 		}
// 	case *StructValue:
// 		nv := n.(*StructValue)
// 		for i := 0; i < ov.NumField(); i++ {
// 			nv.Field(i).SetValue(deepcopy(ov.Field(i), visited))
// 		}

// 	}
// 	return n
// }
