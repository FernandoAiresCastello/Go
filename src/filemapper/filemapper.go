package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

const STANDARD_RIGHTS_ALL = 0x001F0000
const CREATE_NEW = 1
const FILE_ATTRIBUTE_NORMAL = 128

func main() {
	var dll = syscall.NewLazyDLL("kernel32.dll")

	// Create/open a file on disk
	var procCreateFile = dll.NewProc("CreateFileW")
	/*
		https://msdn.microsoft.com/en-us/library/windows/desktop/aa363858(v=vs.85).aspx

		HANDLE WINAPI CreateFile(
		_In_     LPCTSTR               lpFileName,
		_In_     DWORD                 dwDesiredAccess,
		_In_     DWORD                 dwShareMode,
		_In_opt_ LPSECURITY_ATTRIBUTES lpSecurityAttributes,
		_In_     DWORD                 dwCreationDisposition,
		_In_     DWORD                 dwFlagsAndAttributes,
		_In_opt_ HANDLE                hTemplateFile
		);
	*/

	// Close file handle
	var procCloseHandle = dll.NewProc("CloseHandle")
	/*
		https://msdn.microsoft.com/en-us/library/windows/desktop/ms724211(v=vs.85).aspx

		BOOL WINAPI CloseHandle(
		_In_ HANDLE hObject
		);
	*/

	// Map file to memory
	//var procCreateFileMapping = dll.NewProc("CreateFileMappingW")
	/*
		https://msdn.microsoft.com/en-us/library/windows/desktop/aa366537(v=vs.85).aspx

		HANDLE WINAPI CreateFileMapping(
		_In_     HANDLE                hFile,
		_In_opt_ LPSECURITY_ATTRIBUTES lpAttributes,
		_In_     DWORD                 flProtect,
		_In_     DWORD                 dwMaximumSizeHigh,
		_In_     DWORD                 dwMaximumSizeLow,
		_In_opt_ LPCTSTR               lpName
		);
	*/

	fileHandle, _, _ := procCreateFile.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("test.txt"))),
		uintptr(STANDARD_RIGHTS_ALL), 0, 0, CREATE_NEW, FILE_ATTRIBUTE_NORMAL, 0)

	if int(fileHandle) != -1 {
		fmt.Println("File opened/created")

		/*
			ret2, _, _ := procCreateFileMapping.Call(0,
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("This test is Done."))),
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Done Title"))),
			uintptr(MB_YESNOCANCEL))
		*/

		ret, _, _ := procCloseHandle.Call(fileHandle)

		if ret != 0 {
			fmt.Println("File closed")
		}
	} else {
		fmt.Println("Error opening/creating file")
	}
}
