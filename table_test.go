package vinegar

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVigenere_NewTableConfig(t *testing.T) {
	notValidUTF8 := "\xf2a\xf3bc"
	cases := []struct {
		desc                 string // test description
		alphabet             string // input alphabet
		expectedAlphabet     string // formatted alphabet
		keyword              string // input keyword
		expectedKeyword      string // formatted keyword
		secretKey            string // input secretKey
		expectedSecretKey    string // formatted secretKey
		message              string // input message for secret key
		expectedSecretKeyMsg string // formatted secret key for the given message
		fail                 bool   // should the config generation fail
		panic                bool   // should the config generation panic
		err                  string // the expected error string (if any)
	}{
		{
			desc:                 "Success: Standard Latin Alphabet",
			alphabet:             "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			keyword:              "HIDETHETABLE", // contains duplicates
			expectedKeyword:      "HIDETABL",
			expectedAlphabet:     "HIDETABLCFGJKMNOPQRSUVWXYZ",
			secretKey:            "SUPER1SECRET2KEY3", // contains extra characters
			expectedSecretKey:    "SUPERSECRETKEY",
			message:              "THISISAHIDDENMESSAGEISSUPERPRIVATE",
			expectedSecretKeyMsg: "SUPERSECRETKEYSUPERSECRETKEYSUPERS", // matches alphabet length
			fail:                 false,
			panic:                false,
			err:                  "",
		},
		{
			desc:                 "Success: UTF-8 Encoded Alphabet and Keywords",
			alphabet:             "ABC123{}()*&., :世界iqmJKLMAOPJDFÆ»ÆÆ»Æ∏ˆı◊ÇÎ", // contains UTF-8, spaces and punctuation as well as duplicates
			keyword:              "世界MA»Æ»ÆPP∏ˆÌ,.&*)",                          // contains duplicates and characters outside alphabet (Ì)
			expectedKeyword:      "世界MA»ÆP∏ˆ,.&*)",
			expectedAlphabet:     "世界MA»ÆP∏ˆ,.&*)BC123{}( :iqmJKLODFı◊ÇÎ",
			secretKey:            "ÆP∏ˆ,.&*)~`%$|", // contains characters outside of the alphabet
			expectedSecretKey:    "ÆP∏ˆ,.&*)",
			message:              "THIS IS A MESSAGE WITH SPACES, PUNCTUATION AND 世界MA»ÆP∏ SOME UTF-8",
			expectedSecretKeyMsg: "ÆP∏ˆ,.&*)ÆP∏ˆ,.&*)ÆP∏ˆ,.&*)ÆP∏ˆ,.&*)ÆP∏ˆ,.&*)ÆP∏ˆ,.&*)ÆP∏ˆ,.&*)ÆP∏", // same rune count
			fail:                 false,
			panic:                false,
			err:                  "",
		},
		{
			desc:     "Failure: Nil Alphabet",
			alphabet: "",
			fail:     true,
			err:      "Alphabet cannot be empty",
		},
		{
			desc:     "Failure: Nil Secret Key",
			alphabet: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			fail:     true,
			err:      "SecretKey cannot be empty",
		},
		{
			desc:      "Failure: Invalid UTF-8 Alphabet",
			alphabet:  notValidUTF8,
			secretKey: "HIDDEN",
			fail:      true,
			err:       "Invalid alphabet: not UTF-8 encoded",
		},
		{
			desc:      "Failure: Invalid UTF-8 Table Keyword",
			alphabet:  "ABCDEFG",
			secretKey: "HIDDEN",
			keyword:   notValidUTF8,
			fail:      true,
			err:       "Invalid keyword: not UTF-8 encoded",
		},
		{
			desc:      "Failure: Invalid UTF-8 Secret Key",
			alphabet:  "ABCDEFG",
			secretKey: notValidUTF8,
			fail:      true,
			err:       "Invalid secret key: not UTF-8 encoded",
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			config, err := NewTableConfig(c.alphabet, c.keyword, c.secretKey)
			if msg := recover(); msg != "" && c.panic {
				require.Equalf(t,
					msg,
					c.err,
					"Panic message should equal: %s, got: %s\n",
					c.err,
					msg,
				)
				return
			}
			if c.fail {
				require.Error(t, err, "NewTableConfig should fail\n")
				require.EqualError(
					t,
					err,
					c.err,
					"TableConfig error should equal:\n\t'%s'\nGot\n\t'%s'\n",
					c.err,
					err,
				)
				return
			}
			if err != nil {
				require.NoErrorf(
					t, err,
					"NewTableConfig should not have returned an error, got: %s\n",
					err.Error(),
				)
			}
			if config.Validate() != nil {
				require.NoErrorf(
					t,
					config.Validate(),
					"Config Validation Failed: %s\n",
					config.Validate().Error(),
				)
			}
			require.Equalf(
				t,
				string(config.Alphabet),
				c.expectedAlphabet,
				"Expected alphabet:\n\t'%s'\nGot\n\t'%s'\n",
				string(c.expectedAlphabet),
				string(config.Alphabet),
			)
			require.Equalf(
				t,
				string(config.Keyword),
				c.expectedKeyword,
				"Expected keyword:\n\t'%s'\nGot\n\t'%s'\n",
				string(c.expectedKeyword),
				string(config.Keyword),
			)
			require.Equalf(
				t,
				string(config.SecretKey),
				c.expectedSecretKey,
				"Expected secret key:\n\t'%s'\nGot\n\t'%s'\n",
				string(c.expectedSecretKey),
				string(config.SecretKey),
			)
			fmtSecretKey := formatSecretKeyword(
				[]rune(c.secretKey),
				[]rune(c.alphabet),
				c.message,
			)
			require.Equalf(
				t,
				c.expectedSecretKeyMsg,
				string(fmtSecretKey),
				"Expected formatted secret key:\n\t'%s'\nGot\n\t'%s'\n",
				c.expectedSecretKeyMsg,
				string(fmtSecretKey),
			)
		})
	}
}
