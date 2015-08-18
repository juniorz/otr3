// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compat

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"testing"
)

var isQueryTests = []struct {
	msg             string
	expectedVersion int
}{
	{"foo", 0},
	{"?OtR", 0},
	{"?OtR?", 0},
	{"?OTR?", 0},
	{"?OTRv?", 0},
	{"?OTRv1?", 0},
	{"?OTR?v1?", 0},
	{"?OTR?v?", 0},
	{"?OTR?v2?", 2},
	{"?OTRv2?", 2},
	{"?OTRv23?", 2},
	{"?OTRv23 ?", 0},
}

var alicePrivateKeyHex = "000000000080c81c2cb2eb729b7e6fd48e975a932c638b3a9055478583afa46755683e30102447f6da2d8bec9f386bbb5da6403b0040fee8650b6ab2d7f32c55ab017ae9b6aec8c324ab5844784e9a80e194830d548fb7f09a0410df2c4d5c8bc2b3e9ad484e65412be689cf0834694e0839fb2954021521ffdffb8f5c32c14dbf2020b3ce7500000014da4591d58def96de61aea7b04a8405fe1609308d000000808ddd5cb0b9d66956e3dea5a915d9aba9d8a6e7053b74dadb2fc52f9fe4e5bcc487d2305485ed95fed026ad93f06ebb8c9e8baf693b7887132c7ffdd3b0f72f4002ff4ed56583ca7c54458f8c068ca3e8a4dfa309d1dd5d34e2a4b68e6f4338835e5e0fb4317c9e4c7e4806dafda3ef459cd563775a586dd91b1319f72621bf3f00000080b8147e74d8c45e6318c37731b8b33b984a795b3653c2cd1d65cc99efe097cb7eb2fa49569bab5aab6e8a1c261a27d0f7840a5e80b317e6683042b59b6dceca2879c6ffc877a465be690c15e4a42f9a7588e79b10faac11b1ce3741fcef7aba8ce05327a2c16d279ee1b3d77eb783fb10e3356caa25635331e26dd42b8396c4d00000001420bec691fea37ecea58a5c717142f0b804452f57"

var aliceFingerprintHex = "0bb01c360424522e94ee9c346ce877a1a4288b2f"

var bobPrivateKeyHex = "000000000080a5138eb3d3eb9c1d85716faecadb718f87d31aaed1157671d7fee7e488f95e8e0ba60ad449ec732710a7dec5190f7182af2e2f98312d98497221dff160fd68033dd4f3a33b7c078d0d9f66e26847e76ca7447d4bab35486045090572863d9e4454777f24d6706f63e02548dfec2d0a620af37bbc1d24f884708a212c343b480d00000014e9c58f0ea21a5e4dfd9f44b6a9f7f6a9961a8fa9000000803c4d111aebd62d3c50c2889d420a32cdf1e98b70affcc1fcf44d59cca2eb019f6b774ef88153fb9b9615441a5fe25ea2d11b74ce922ca0232bd81b3c0fcac2a95b20cb6e6c0c5c1ace2e26f65dc43c751af0edbb10d669890e8ab6beea91410b8b2187af1a8347627a06ecea7e0f772c28aae9461301e83884860c9b656c722f0000008065af8625a555ea0e008cd04743671a3cda21162e83af045725db2eb2bb52712708dc0cc1a84c08b3649b88a966974bde27d8612c2861792ec9f08786a246fcadd6d8d3a81a32287745f309238f47618c2bd7612cb8b02d940571e0f30b96420bcd462ff542901b46109b1e5ad6423744448d20a57818a8cbb1647d0fea3b664e0000001440f9f2eb554cb00d45a5826b54bfa419b6980e48"

func TestKeySerialization(t *testing.T) {
	var priv PrivateKey
	alicePrivateKey, _ := hex.DecodeString(alicePrivateKeyHex)
	rest, ok := priv.Parse(alicePrivateKey)
	if !ok {
		t.Error("failed to parse private key")
	}
	if len(rest) > 0 {
		t.Error("data remaining after parsing private key")
	}

	out := priv.Serialize(nil)
	if !bytes.Equal(alicePrivateKey, out) {
		t.Errorf("serialization (%x) is not equal to original (%x)", out, alicePrivateKey)
	}

	aliceFingerprint, _ := hex.DecodeString(aliceFingerprintHex)
	fingerprint := priv.PublicKey.Fingerprint()
	if !bytes.Equal(aliceFingerprint, fingerprint) {
		t.Errorf("fingerprint (%x) is not equal to expected value (%x)", fingerprint, aliceFingerprint)
	}
}

const libOTRPrivateKey = `(privkeys
 (account
(name "foo@example.com")
(protocol prpl-jabber)
(private-key
 (dsa
  (p #00FC07ABCF0DC916AFF6E9AE47BEF60C7AB9B4D6B2469E436630E36F8A489BE812486A09F30B71224508654940A835301ACC525A4FF133FC152CC53DCC59D65C30A54F1993FE13FE63E5823D4C746DB21B90F9B9C00B49EC7404AB1D929BA7FBA12F2E45C6E0A651689750E8528AB8C031D3561FECEE72EBB4A090D450A9B7A857#)
  (q #00997BD266EF7B1F60A5C23F3A741F2AEFD07A2081#)
  (g #535E360E8A95EBA46A4F7DE50AD6E9B2A6DB785A66B64EB9F20338D2A3E8FB0E94725848F1AA6CC567CB83A1CC517EC806F2E92EAE71457E80B2210A189B91250779434B41FC8A8873F6DB94BEA7D177F5D59E7E114EE10A49CFD9CEF88AE43387023B672927BA74B04EB6BBB5E57597766A2F9CE3857D7ACE3E1E3BC1FC6F26#)
  (y #0AC8670AD767D7A8D9D14CC1AC6744CD7D76F993B77FFD9E39DF01E5A6536EF65E775FCEF2A983E2A19BD6415500F6979715D9FD1257E1FE2B6F5E1E74B333079E7C880D39868462A93454B41877BE62E5EF0A041C2EE9C9E76BD1E12AE25D9628DECB097025DD625EF49C3258A1A3C0FF501E3DC673B76D7BABF349009B6ECF#)
  (x #14D0345A3562C480A039E3C72764F72D79043216#)
  )
 )
 )
)`

func TestParseLibOTRPrivateKey(t *testing.T) {
	var priv PrivateKey

	if !priv.Import([]byte(libOTRPrivateKey)) {
		t.Fatalf("Failed to import sample private key")
	}
}

func TestSignVerify(t *testing.T) {
	var priv PrivateKey
	alicePrivateKey, _ := hex.DecodeString(alicePrivateKeyHex)
	_, ok := priv.Parse(alicePrivateKey)
	if !ok {
		t.Error("failed to parse private key")
	}

	var msg [32]byte
	rand.Reader.Read(msg[:])

	sig := priv.Sign(rand.Reader, msg[:])
	rest, ok := priv.PublicKey.Verify(msg[:], sig)
	if !ok {
		t.Errorf("signature (%x) of %x failed to verify", sig, msg[:])
	} else if len(rest) > 0 {
		t.Error("signature data remains after verification")
	}

	sig[10] ^= 80
	_, ok = priv.PublicKey.Verify(msg[:], sig)
	if ok {
		t.Errorf("corrupted signature (%x) of %x verified", sig, msg[:])
	}
}

func TestConversation(t *testing.T) {
	alicePrivateKey, _ := hex.DecodeString(alicePrivateKeyHex)
	bobPrivateKey, _ := hex.DecodeString(bobPrivateKeyHex)

	var alice, bob Conversation
	alice.PrivateKey = new(PrivateKey)
	bob.PrivateKey = new(PrivateKey)
	alice.PrivateKey.Parse(alicePrivateKey)
	bob.PrivateKey.Parse(bobPrivateKey)
	alice.FragmentSize = 100
	bob.FragmentSize = 100

	var alicesMessage, bobsMessage [][]byte
	var out []byte
	var aliceChange, bobChange SecurityChange
	var err error
	alicesMessage = append(alicesMessage, []byte(QueryMessage))

	if alice.IsEncrypted() {
		t.Error("Alice believes that the conversation is secure before we've started")
	}
	if bob.IsEncrypted() {
		t.Error("Bob believes that the conversation is secure before we've started")
	}

	for round := 0; len(alicesMessage) > 0 || len(bobsMessage) > 0; round++ {
		bobsMessage = nil
		for i, msg := range alicesMessage {
			out, _, bobChange, bobsMessage, err = bob.Receive(msg)
			if len(out) > 0 {
				t.Errorf("Bob generated output during key exchange, round %d, message %d", round, i)
			}
			if err != nil {
				t.Fatalf("Bob returned an error, round %d, message %d (%x): %s", round, i, msg, err)
			}
			if len(bobsMessage) > 0 && i != len(alicesMessage)-1 {
				t.Errorf("Bob produced output while processing a fragment, round %d, message %d", round, i)
			}
		}

		alicesMessage = nil
		for i, msg := range bobsMessage {
			out, _, aliceChange, alicesMessage, err = alice.Receive(msg)
			if len(out) > 0 {
				t.Errorf("Alice generated output during key exchange, round %d, message %d", round, i)
			}
			if err != nil {
				t.Fatalf("Alice returned an error, round %d, message %d (%x): %s", round, i, msg, err)
			}
			if len(alicesMessage) > 0 && i != len(bobsMessage)-1 {
				t.Errorf("Alice produced output while processing a fragment, round %d, message %d", round, i)
			}
		}
	}

	if aliceChange != NewKeys {
		t.Errorf("Alice terminated without signaling new keys")
	}
	if bobChange != NewKeys {
		t.Errorf("Bob terminated without signaling new keys")
	}

	if !bytes.Equal(alice.SSID[:], bob.SSID[:]) {
		t.Errorf("Session identifiers don't match. Alice has %x, Bob has %x", alice.SSID[:], bob.SSID[:])
	}

	if !alice.IsEncrypted() {
		t.Error("Alice doesn't believe that the conversation is secure")
	}
	if !bob.IsEncrypted() {
		t.Error("Bob doesn't believe that the conversation is secure")
	}

	var testMessage = []byte("hello Bob")
	alicesMessage, err = alice.Send(testMessage)
	for i, msg := range alicesMessage {
		out, encrypted, _, _, err := bob.Receive(msg)
		if err != nil {
			t.Errorf("Error generated while processing test message: %s", err.Error())
		}
		if len(out) > 0 {
			if i != len(alicesMessage)-1 {
				t.Fatal("Bob produced a message while processing a fragment of Alice's")
			}
			if !encrypted {
				t.Errorf("Message was not marked as encrypted")
			}
			if !bytes.Equal(out, testMessage) {
				t.Errorf("Message corrupted: got %x, want %x", out, testMessage)
			}
		}
	}
}
