package kontrol

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/koding/kite"
	"github.com/koding/kite/protocol"
	"github.com/koding/kite/testkeys"
)

func testinsie(t *testing.T) {
	kont, conf := startKontrol(testkeys.PrivateThird, testkeys.PublicThird, 5501)
	defer kont.Close()
	// kont.HandleWebRTC(r *kite.Request)
	hk1, err := NewHelloKite("kite1", conf)
	if err != nil {
		t.Fatalf("error creating kite1: %s", err)
	}
	defer hk1.Close()
	hk2, err := NewHelloKite("kite2", conf)
	if err != nil {
		t.Fatalf("error creating kite1: %s", err)
	}
	defer hk2.Close()

	err = hk1.Kite.SendWebRTCRequest(&protocol.WebRTCSignalMessage{Dst: hk2.Kite.Id})
	fmt.Println("hk1 err-->", err)

	err = hk2.Kite.SendWebRTCRequest(&protocol.WebRTCSignalMessage{Dst: hk1.Kite.Id})
	fmt.Println("hk2 err-->", err)

	hk2.Token = true
}
func TestKontrol_HandleWebRTC(t *testing.T) {
	fmt.Println("2-->", 2)
	testinsie(t)
	fmt.Println("1-->", 1)
	tests := []struct {
		name    string
		kontrol *Kontrol
		args    *kite.Request
		want    interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// k := &Kontrol{}
			got, err := tt.kontrol.Kite.WebRTCHandler(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Kontrol.HandleWebRTC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Kontrol.HandleWebRTC() = %v, want %v", got, tt.want)
			}
		})
	}
}
