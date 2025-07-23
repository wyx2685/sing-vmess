//go:build with_utls

package vless

import (
	"net"
	"reflect"
	"unsafe"

	"github.com/sagernet/sing/common"

	"github.com/xtls/reality"
)

func init() {
	tlsRegistry = append(tlsRegistry, func(conn net.Conn) (loaded bool, netConn net.Conn, reflectType reflect.Type, reflectPointer uintptr) {
		tlsConn, loaded := common.Cast[*reality.Conn](conn)
		if loaded {
			return true, tlsConn.NetConn(), reflect.TypeOf(tlsConn).Elem(), uintptr(unsafe.Pointer(tlsConn))
		}
		return
	})
}
