package main

import (
	"VMProtect"
	"time"
)

func main() {
	VMProtect.BeginUltra("Marker\x00")
	str := VMProtect.GoString(VMProtect.DecryptStringA("This is a decrypted string\x00"))
	serial := "SerialNumber\x00"

	println(str)
	println("HWID: ", VMProtect.GetCurrentHWID())
	println("IsProtected: ", VMProtect.IsProtected())
	println("IsDebuggerPresent: ", VMProtect.IsDebuggerPresent(true))
	println("IsVirtualMachinePresent: ", VMProtect.IsVirtualMachinePresent())
	println("IsValidImageCRC: ", VMProtect.IsValidImageCRC())
	println("SetSerialNumber: ", VMProtect.SetSerialNumber(serial))
	if VMProtect.GetSerialNumberState() == VMProtect.SERIAL_STATE_SUCCESS {
		println("-- Registered --")
	}
	println("User: ", VMProtect.GetUser())
	println("Email: ", VMProtect.GetEmail())
	println("ExpireDate: ", VMProtect.GetExpireDate())
	println("MaxBuildDate: ", VMProtect.GetMaxBuild())
	println("RunningTimeLimit: ", VMProtect.GetRunningTimeLimit())
	time.Sleep(3 * time.Minute)
	if VMProtect.GetSerialNumberState() == VMProtect.SERIAL_STATE_FLAG_RUNNING_TIME_OVER {
		println("-- Running Time Over, Please Registere --")
	}
	VMProtect.End()
}
