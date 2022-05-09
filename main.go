package main

import (
	"fmt"
	"syscall"
	"time"
)

const (
	_NULL               = 0x00
	_KEYEVENTF_KEYUP    = 0x0002
	_KEYEVENTF_SCANCODE = 0x0008
)

var dll = syscall.NewLazyDLL("user32.dll")
var procKeyBd = dll.NewProc("keybd_event")

func main() {
	for {
		hourNow := time.Now().Hour()

		if hourNow >= 9 && hourNow <= 16 {

			err := pressKey(_NULL)
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(1 * time.Minute)
		}
	}
}

// Launch key bounding
func pressKey(key int) error {
	err := downKey(key)
	if err != nil {
		return err
	}
	err = upKey(key)
	if err != nil {
		return err
	}

	return nil
}

func downKey(key int) error {
	flag := 0
	if key < 0xFFF { // Detect if the key code is virtual or no
		flag |= _KEYEVENTF_SCANCODE
	} else {
		key -= 0xFFF
	}
	vkey := key + 0x80
	_, _, err := procKeyBd.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
	return checkError(err)
}

func upKey(key int) error {
	flag := _KEYEVENTF_KEYUP
	if key < 0xFFF {
		flag |= _KEYEVENTF_SCANCODE
	} else {
		key -= 0xFFF
	}
	vkey := key + 0x80
	_, _, err := procKeyBd.Call(uintptr(key), uintptr(vkey), uintptr(flag), 0)
	return checkError(err)
}

func checkError(err error) error {
	e, ok := err.(syscall.Errno)
	if ok && e == 0 {
		return nil
	}
	return err
}
