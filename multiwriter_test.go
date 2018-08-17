package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/ugorji/go/codec"
	"golang.org/x/crypto/blake2b"
)

// msgpackHandle is a shared handle for encoding/decoding of structs
var MsgpackHandle = func() *codec.MsgpackHandle {
	h := &codec.MsgpackHandle{RawToString: true}

	// Sets the default type for decoding a map into a nil interface{}.
	// This is necessary in particular because we store the driver configs as a
	// nil interface{}.
	h.MapType = reflect.TypeOf(map[string]interface{}(nil))
	return h
}()

func getPayload() map[string]interface{} {
	keys, _ := strconv.Atoi(os.Getenv("N"))
	if keys == 0 {
		keys = 100
	}
	p := make(map[string]interface{}, keys)
	for i := 0; i < keys; i++ {
		p[fmt.Sprintf("%x", i)] = "0123456789"
	}
	return p
}

func BenchmarkJSON_Blake2b_Multi(b *testing.B) {
	payload := getPayload()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h, _ := blake2b.New256(nil)
		var buf bytes.Buffer
		w := io.MultiWriter(h, &buf)

		if err := json.NewEncoder(w).Encode(payload); err != nil {
			b.Fatalf("err: %v", err)
		}

		b.SetBytes(int64(buf.Len() + len(h.Sum(nil))))
	}
}

func BenchmarkJSON_Blake2b_NoMulti(b *testing.B) {
	payload := getPayload()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(payload); err != nil {
			b.Fatalf("err: %v", err)
		}

		b.SetBytes(int64(buf.Len() + len(blake2b.Sum256(buf.Bytes()))))
	}
}

func BenchmarkJSON_MD5_Multi(b *testing.B) {
	payload := getPayload()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := md5.New()
		var buf bytes.Buffer
		w := io.MultiWriter(h, &buf)

		if err := json.NewEncoder(w).Encode(payload); err != nil {
			b.Fatalf("err: %v", err)
		}

		b.SetBytes(int64(buf.Len() + len(h.Sum(nil))))
	}
}

func BenchmarkJSON_MD5_NoMulti(b *testing.B) {
	payload := getPayload()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(payload); err != nil {
			b.Fatalf("err: %v", err)
		}

		b.SetBytes(int64(buf.Len() + len(md5.Sum(buf.Bytes()))))
	}
}

func BenchmarkMsgpack_Blake2b_Multi(b *testing.B) {
	payload := getPayload()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h, _ := blake2b.New256(nil)
		var buf bytes.Buffer
		w := io.MultiWriter(h, &buf)

		if err := codec.NewEncoder(w, MsgpackHandle).Encode(payload); err != nil {
			b.Fatalf("err: %v", err)
		}

		b.SetBytes(int64(buf.Len() + len(h.Sum(nil))))
	}
}

func BenchmarkMsgpack_Blake2b_NoMulti(b *testing.B) {
	payload := getPayload()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := codec.NewEncoder(&buf, MsgpackHandle).Encode(payload); err != nil {
			b.Fatalf("err: %v", err)
		}

		b.SetBytes(int64(buf.Len() + len(blake2b.Sum256(buf.Bytes()))))
	}
}

func BenchmarkMsgpack_MD5_Multi(b *testing.B) {
	payload := getPayload()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := md5.New()
		var buf bytes.Buffer
		w := io.MultiWriter(h, &buf)

		if err := codec.NewEncoder(w, MsgpackHandle).Encode(payload); err != nil {
			b.Fatalf("err: %v", err)
		}

		b.SetBytes(int64(buf.Len() + len(h.Sum(nil))))
	}
}

func BenchmarkMsgpack_MD5_NoMulti(b *testing.B) {
	payload := getPayload()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		if err := codec.NewEncoder(&buf, MsgpackHandle).Encode(payload); err != nil {
			b.Fatalf("err: %v", err)
		}

		b.SetBytes(int64(buf.Len() + len(md5.Sum(buf.Bytes()))))
	}
}
