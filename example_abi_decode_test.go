package eos_test

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	eos "github.com/eoscanada/eos-go"
)

func Test_ExampleABI_DecodeTableRowTyped(t *testing.T) {
	abi, err := eos.NewABI(strings.NewReader(abiJSON()))
	if err != nil {
		panic(fmt.Errorf("get ABI: %w", err))
	}

	tableDef := abi.TableForName(eos.TableName("variant"))
	if tableDef == nil {
		panic(fmt.Errorf("table be should be present"))
	}

	bytes, err := abi.DecodeTableRowTyped(tableDef.Type, data())
	if err != nil {
		panic(fmt.Errorf("decode row: %w", err))
	}

	fmt.Println(string(bytes))
}

func data() []byte {
	bytes, err := hex.DecodeString(`00000000000000000c00000000ffffffff000000`)
	if err != nil {
		panic(fmt.Errorf("decode data: %w", err))
	}

	return bytes
}

func abiJSON() string {
	return `
		{
  "version": "eosio::abi/1.1",
  "types": [
    {
      "new_type_name": "varying_action",
      "type": "variant_uint16_string"
    }
  ],
  "structs": [
    {
      "name": "creaorder",
      "base": "",
      "fields": [
        {
          "name": "n1",
          "type": "name"
        },
        {
          "name": "n2",
          "type": "name"
        },
        {
          "name": "n3",
          "type": "name"
        },
        {
          "name": "n4",
          "type": "name"
        },
        {
          "name": "n5",
          "type": "name"
        }
      ]
    },
    {
      "name": "dbins",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        }
      ]
    },
    {
      "name": "dbinstwo",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "first",
          "type": "uint64"
        },
        {
          "name": "second",
          "type": "uint64"
        }
      ]
    },
    {
      "name": "dbrem",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        }
      ]
    },
    {
      "name": "dbremtwo",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "first",
          "type": "uint64"
        },
        {
          "name": "second",
          "type": "uint64"
        }
      ]
    },
    {
      "name": "dbupd",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        }
      ]
    },
    {
      "name": "dtrx",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "fail_now",
          "type": "bool"
        },
        {
          "name": "fail_later",
          "type": "bool"
        },
        {
          "name": "fail_later_nested",
          "type": "bool"
        },
        {
          "name": "delay_sec",
          "type": "uint32"
        },
        {
          "name": "nonce",
          "type": "string"
        }
      ]
    },
    {
      "name": "dtrxcancel",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        }
      ]
    },
    {
      "name": "dtrxexec",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "fail",
          "type": "bool"
        },
        {
          "name": "failNested",
          "type": "bool"
        },
        {
          "name": "nonce",
          "type": "string"
        }
      ]
    },
    {
      "name": "inlinedeep",
      "base": "",
      "fields": [
        {
          "name": "tag",
          "type": "string"
        },
        {
          "name": "n4",
          "type": "name"
        },
        {
          "name": "n5",
          "type": "name"
        },
        {
          "name": "nestedInlineTag",
          "type": "string"
        },
        {
          "name": "nestedInlineFail",
          "type": "bool"
        },
        {
          "name": "nestedCfaInlineTag",
          "type": "string"
        }
      ]
    },
    {
      "name": "inlineempty",
      "base": "",
      "fields": [
        {
          "name": "tag",
          "type": "string"
        },
        {
          "name": "fail",
          "type": "bool"
        }
      ]
    },
    {
      "name": "member_row",
      "base": "",
      "fields": [
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "amount",
          "type": "asset"
        },
        {
          "name": "memo",
          "type": "string"
        },
        {
          "name": "created_at",
          "type": "time_point_sec"
        },
        {
          "name": "expires_at",
          "type": "time_point_sec"
        }
      ]
    },
    {
      "name": "nestdtrxexec",
      "base": "",
      "fields": [
        {
          "name": "fail",
          "type": "bool"
        }
      ]
    },
    {
      "name": "nestonerror",
      "base": "",
      "fields": [
        {
          "name": "fail",
          "type": "bool"
        }
      ]
    },
    {
      "name": "producerows",
      "base": "",
      "fields": [
        {
          "name": "row_count",
          "type": "uint64"
        }
      ]
    },
    {
      "name": "sk_row",
      "base": "",
      "fields": [
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "i64",
          "type": "uint64"
        },
        {
          "name": "i128",
          "type": "uint128"
        },
        {
          "name": "d64",
          "type": "float64"
        },
        {
          "name": "d128",
          "type": "float128"
        },
        {
          "name": "c256",
          "type": "checksum256"
        },
        {
          "name": "unrelated",
          "type": "uint64"
        }
      ]
    },
    {
      "name": "sktest",
      "base": "",
      "fields": [
        {
          "name": "action",
          "type": "name"
        }
      ]
    },
    {
      "name": "variant_row",
      "base": "",
      "fields": [
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "variant_field",
          "type": "variant_int8_uint16_uint32_int32"
        },
        {
          "name": "creation_number",
          "type": "uint64"
        }
      ]
    },
    {
      "name": "varianttest",
      "base": "",
      "fields": [
        {
          "name": "value",
          "type": "varying_action"
        }
      ]
    }
  ],
  "actions": [
    {
      "name": "creaorder",
      "type": "creaorder",
      "ricardian_contract": ""
    },
    {
      "name": "dbins",
      "type": "dbins",
      "ricardian_contract": ""
    },
    {
      "name": "dbinstwo",
      "type": "dbinstwo",
      "ricardian_contract": ""
    },
    {
      "name": "dbrem",
      "type": "dbrem",
      "ricardian_contract": ""
    },
    {
      "name": "dbremtwo",
      "type": "dbremtwo",
      "ricardian_contract": ""
    },
    {
      "name": "dbupd",
      "type": "dbupd",
      "ricardian_contract": ""
    },
    {
      "name": "dtrx",
      "type": "dtrx",
      "ricardian_contract": ""
    },
    {
      "name": "dtrxcancel",
      "type": "dtrxcancel",
      "ricardian_contract": ""
    },
    {
      "name": "dtrxexec",
      "type": "dtrxexec",
      "ricardian_contract": ""
    },
    {
      "name": "inlinedeep",
      "type": "inlinedeep",
      "ricardian_contract": ""
    },
    {
      "name": "inlineempty",
      "type": "inlineempty",
      "ricardian_contract": ""
    },
    {
      "name": "nestdtrxexec",
      "type": "nestdtrxexec",
      "ricardian_contract": ""
    },
    {
      "name": "nestonerror",
      "type": "nestonerror",
      "ricardian_contract": ""
    },
    {
      "name": "producerows",
      "type": "producerows",
      "ricardian_contract": ""
    },
    {
      "name": "sktest",
      "type": "sktest",
      "ricardian_contract": ""
    },
    {
      "name": "varianttest",
      "type": "varianttest",
      "ricardian_contract": ""
    }
  ],
  "tables": [
    {
      "name": "member",
      "index_type": "i64",
      "type": "member_row"
    },
    {
      "name": "sk.c",
      "index_type": "i64",
      "type": "sk_row"
    },
    {
      "name": "sk.d",
      "index_type": "i64",
      "type": "sk_row"
    },
    {
      "name": "sk.dd",
      "index_type": "i64",
      "type": "sk_row"
    },
    {
      "name": "sk.i",
      "index_type": "i64",
      "type": "sk_row"
    },
    {
      "name": "sk.ii",
      "index_type": "i64",
      "type": "sk_row"
    },
    {
      "name": "sk.multi",
      "index_type": "i64",
      "type": "sk_row"
    },
    {
      "name": "variant",
      "index_type": "i64",
      "type": "variant_row"
    }
  ],
  "variants": [
    {
      "name": "variant_int8_uint16_uint32_int32",
      "types": [
        "int8",
        "uint16",
        "uint32",
        "int32"
      ]
    },
    {
      "name": "variant_uint16_string",
      "types": [
        "uint16",
        "string"
      ]
    }
  ]
}
`
}
