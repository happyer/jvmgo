package reserved

import "ch11/instructions/base"
import "ch11/rtda"
import "ch11/native"
import _ "ch11/native/java/io"
import _ "ch11/native/java/lang"
import _ "ch11/native/java/security"
import _ "ch11/native/java/util/concurrent/atomic"
import _ "ch11/native/sun/io"
import _ "ch11/native/sun/misc"
import _ "ch11/native/sun/reflect"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
