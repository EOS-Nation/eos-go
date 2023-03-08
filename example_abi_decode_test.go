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

	tableDef := abi.TableForName(eos.TableName("account"))
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
	bytes, err := hex.DecodeString(`00000000005c95b191025917d60000000004454f5300000000`)
	if err != nil {
		panic(fmt.Errorf("decode data: %w", err))
	}

	return bytes
}

func abiJSON() string {
	return `
{
  "version": "eosio::abi/1.3",
  "types": [
    {
      "new_type_name": "INT8_VEC",
      "type": "int8[]"
    },
    {
      "new_type_name": "INT16_VEC",
      "type": "int16[]"
    },
    {
      "new_type_name": "INT32_VEC",
      "type": "int32[]"
    },
    {
      "new_type_name": "INT64_VEC",
      "type": "int64[]"
    },
    {
      "new_type_name": "UINT8_VEC",
      "type": "uint8[]"
    },
    {
      "new_type_name": "UINT16_VEC",
      "type": "uint16[]"
    },
    {
      "new_type_name": "UINT32_VEC",
      "type": "uint32[]"
    },
    {
      "new_type_name": "UINT64_VEC",
      "type": "uint64[]"
    },
    {
      "new_type_name": "FLOAT_VEC",
      "type": "float32[]"
    },
    {
      "new_type_name": "DOUBLE_VEC",
      "type": "float64[]"
    },
    {
      "new_type_name": "STRING_VEC",
      "type": "string[]"
    }
  ],
  "structs": [
    {
      "name": "no_auth",
      "base": ""
    },
    {
      "name": "account_auth",
      "base": "",
      "fields": [
        {
          "name": "contract",
          "type": "name"
        },
        {
          "name": "contract_account",
          "type": "name"
        },
        {
          "name": "eosio_account",
          "type": "name"
        }
      ]
    },
    {
      "name": "signature_auth",
      "base": "",
      "fields": [
        {
          "name": "signature",
          "type": "signature"
        },
        {
          "name": "contract",
          "type": "name"
        },
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "sequence",
          "type": "varuint32"
        }
      ]
    },
    {
      "name": "withdraw",
      "base": "",
      "fields": [
        {
          "name": "owner",
          "type": "name"
        },
        {
          "name": "quantity",
          "type": "asset"
        }
      ]
    },
    {
      "name": "donate",
      "base": "",
      "fields": [
        {
          "name": "owner",
          "type": "name"
        },
        {
          "name": "quantity",
          "type": "asset"
        }
      ]
    },
    {
      "name": "fundtransfer",
      "base": "",
      "fields": [
        {
          "name": "from",
          "type": "name"
        },
        {
          "name": "distribution_time",
          "type": "block_timestamp_type"
        },
        {
          "name": "rank",
          "type": "uint8"
        },
        {
          "name": "to",
          "type": "name"
        },
        {
          "name": "amount",
          "type": "asset"
        },
        {
          "name": "memo",
          "type": "string"
        }
      ]
    },
    {
      "name": "usertransfer",
      "base": "",
      "fields": [
        {
          "name": "from",
          "type": "name"
        },
        {
          "name": "to",
          "type": "name"
        },
        {
          "name": "amount",
          "type": "asset"
        },
        {
          "name": "memo",
          "type": "string"
        }
      ]
    },
    {
      "name": "attribute",
      "base": "",
      "fields": [
        {
          "name": "key",
          "type": "string"
        },
        {
          "name": "value",
          "type": "attribute_value"
        }
      ]
    },
    {
      "name": "genesis",
      "base": "",
      "fields": [
        {
          "name": "community",
          "type": "string"
        },
        {
          "name": "community_symbol",
          "type": "symbol"
        },
        {
          "name": "minimum_donation",
          "type": "asset"
        },
        {
          "name": "initial_members",
          "type": "name[]"
        },
        {
          "name": "genesis_video",
          "type": "string"
        },
        {
          "name": "collection_attributes",
          "type": "attribute[]"
        },
        {
          "name": "auction_starting_bid",
          "type": "asset"
        },
        {
          "name": "auction_duration",
          "type": "uint32"
        },
        {
          "name": "memo",
          "type": "string"
        },
        {
          "name": "election_day",
          "type": "uint8"
        },
        {
          "name": "election_time",
          "type": "string"
        }
      ]
    },
    {
      "name": "setmindonfee",
      "base": "",
      "fields": [
        {
          "name": "new_minimum_donation",
          "type": "asset"
        }
      ]
    },
    {
      "name": "addtogenesis",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "expiration",
          "type": "time_point"
        }
      ]
    },
    {
      "name": "gensetexpire",
      "base": "",
      "fields": [
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "new_expiration",
          "type": "time_point"
        }
      ]
    },
    {
      "name": "clearall",
      "base": ""
    },
    {
      "name": "inductinit",
      "base": "",
      "fields": [
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "inviter",
          "type": "name"
        },
        {
          "name": "invitee",
          "type": "name"
        },
        {
          "name": "witnesses",
          "type": "name[]"
        }
      ]
    },
    {
      "name": "encrypted_key",
      "base": "",
      "fields": [
        {
          "name": "sender_key",
          "type": "public_key"
        },
        {
          "name": "recipient_key",
          "type": "public_key"
        },
        {
          "name": "key",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "inductmeetin",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "keys",
          "type": "encrypted_key[]"
        },
        {
          "name": "data",
          "type": "bytes"
        },
        {
          "name": "old_data",
          "type": "bytes?"
        }
      ]
    },
    {
      "name": "new_member_profile",
      "base": "",
      "fields": [
        {
          "name": "name",
          "type": "string"
        },
        {
          "name": "img",
          "type": "string"
        },
        {
          "name": "bio",
          "type": "string"
        },
        {
          "name": "social",
          "type": "string"
        },
        {
          "name": "attributions",
          "type": "string"
        }
      ]
    },
    {
      "name": "inductprofil",
      "base": "",
      "fields": [
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "new_member_profile",
          "type": "new_member_profile"
        }
      ]
    },
    {
      "name": "inductvideo",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "video",
          "type": "string"
        }
      ]
    },
    {
      "name": "inductendors",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "induction_data_hash",
          "type": "checksum256"
        }
      ]
    },
    {
      "name": "setencpubkey",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "key",
          "type": "public_key"
        }
      ]
    },
    {
      "name": "electsettime",
      "base": "",
      "fields": [
        {
          "name": "election_time",
          "type": "time_point_sec"
        }
      ]
    },
    {
      "name": "electconfig",
      "base": "",
      "fields": [
        {
          "name": "day",
          "type": "uint8"
        },
        {
          "name": "time",
          "type": "string"
        },
        {
          "name": "round_duration",
          "type": "uint32"
        }
      ]
    },
    {
      "name": "electopt",
      "base": "",
      "fields": [
        {
          "name": "member",
          "type": "name"
        },
        {
          "name": "participating",
          "type": "bool"
        }
      ]
    },
    {
      "name": "electseed",
      "base": "",
      "fields": [
        {
          "name": "btc_header",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "electmeeting",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "round",
          "type": "uint8"
        },
        {
          "name": "keys",
          "type": "encrypted_key[]"
        },
        {
          "name": "data",
          "type": "bytes"
        },
        {
          "name": "old_data",
          "type": "bytes?"
        }
      ]
    },
    {
      "name": "electvote",
      "base": "",
      "fields": [
        {
          "name": "round",
          "type": "uint8"
        },
        {
          "name": "voter",
          "type": "name"
        },
        {
          "name": "candidate",
          "type": "name"
        }
      ]
    },
    {
      "name": "electvideo",
      "base": "",
      "fields": [
        {
          "name": "round",
          "type": "uint8"
        },
        {
          "name": "voter",
          "type": "name"
        },
        {
          "name": "video",
          "type": "string"
        }
      ]
    },
    {
      "name": "electprocess",
      "base": "",
      "fields": [
        {
          "name": "max_steps",
          "type": "uint32"
        }
      ]
    },
    {
      "name": "bylawspropose",
      "base": "",
      "fields": [
        {
          "name": "proposer",
          "type": "name"
        },
        {
          "name": "bylaws",
          "type": "string"
        }
      ]
    },
    {
      "name": "bylawsapprove",
      "base": "",
      "fields": [
        {
          "name": "approver",
          "type": "name"
        },
        {
          "name": "bylaws_hash",
          "type": "checksum256"
        }
      ]
    },
    {
      "name": "bylawsratify",
      "base": "",
      "fields": [
        {
          "name": "approver",
          "type": "name"
        },
        {
          "name": "bylaws_hash",
          "type": "checksum256"
        }
      ]
    },
    {
      "name": "distribute",
      "base": "",
      "fields": [
        {
          "name": "max_steps",
          "type": "uint32"
        }
      ]
    },
    {
      "name": "givesbt",
      "base": "",
      "fields": [
        {
          "name": "max_steps",
          "type": "uint32"
        }
      ]
    },
    {
      "name": "setdistpct",
      "base": "",
      "fields": [
        {
          "name": "pct",
          "type": "uint8"
        }
      ]
    },
    {
      "name": "inductdonate",
      "base": "",
      "fields": [
        {
          "name": "payer",
          "type": "name"
        },
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "quantity",
          "type": "asset"
        }
      ]
    },
    {
      "name": "inductcancel",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "id",
          "type": "uint64"
        }
      ]
    },
    {
      "name": "inducted",
      "base": "",
      "fields": [
        {
          "name": "inductee",
          "type": "name"
        }
      ]
    },
    {
      "name": "resign",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        }
      ]
    },
    {
      "name": "removemember",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "memo",
          "type": "string"
        }
      ]
    },
    {
      "name": "rename",
      "base": "",
      "fields": [
        {
          "name": "old_account",
          "type": "name"
        },
        {
          "name": "new_account",
          "type": "name"
        }
      ]
    },
    {
      "name": "gc",
      "base": "",
      "fields": [
        {
          "name": "limit",
          "type": "uint32"
        }
      ]
    },
    {
      "name": "migrate",
      "base": "",
      "fields": [
        {
          "name": "limit",
          "type": "uint32"
        }
      ]
    },
    {
      "name": "unmigrate",
      "base": ""
    },
    {
      "name": "unsupported_verb",
      "base": ""
    },
    {
      "name": "account_v0",
      "base": "",
      "fields": [
        {
          "name": "owner",
          "type": "name"
        },
        {
          "name": "balance",
          "type": "asset"
        }
      ]
    },
    {
      "name": "auction_v0",
      "base": "",
      "fields": [
        {
          "name": "asset_id",
          "type": "uint64"
        },
        {
          "name": "last_known_end_time",
          "type": "time_point_sec"
        }
      ]
    },
    {
      "name": "bylaws_v0",
      "base": "",
      "fields": [
        {
          "name": "type",
          "type": "name"
        },
        {
          "name": "text",
          "type": "string"
        },
        {
          "name": "time",
          "type": "block_timestamp_type"
        },
        {
          "name": "approvals",
          "type": "name[]"
        }
      ]
    },
    {
      "name": "distribution_account_v0",
      "base": "",
      "fields": [
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "owner",
          "type": "name"
        },
        {
          "name": "distribution_time",
          "type": "block_timestamp_type"
        },
        {
          "name": "rank",
          "type": "uint8"
        },
        {
          "name": "balance",
          "type": "asset"
        }
      ]
    },
    {
      "name": "next_distribution",
      "base": "",
      "fields": [
        {
          "name": "distribution_time",
          "type": "block_timestamp_type"
        }
      ]
    },
    {
      "name": "election_distribution",
      "base": "",
      "fields": [
        {
          "name": "distribution_time",
          "type": "block_timestamp_type"
        },
        {
          "name": "amount",
          "type": "asset"
        }
      ]
    },
    {
      "name": "current_distribution",
      "base": "",
      "fields": [
        {
          "name": "distribution_time",
          "type": "block_timestamp_type"
        },
        {
          "name": "last_processed",
          "type": "name"
        },
        {
          "name": "rank_distribution",
          "type": "asset[]"
        }
      ]
    },
    {
      "name": "badge_v0",
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
          "name": "round",
          "type": "uint8"
        },
        {
          "name": "vote_time",
          "type": "time_point"
        },
        {
          "name": "end_vote_time",
          "type": "time_point"
        }
      ]
    },
    {
      "name": "endorsement_v0",
      "base": "",
      "fields": [
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "inviter",
          "type": "name"
        },
        {
          "name": "invitee",
          "type": "name"
        },
        {
          "name": "endorser",
          "type": "name"
        },
        {
          "name": "induction_id",
          "type": "uint64"
        },
        {
          "name": "endorsed",
          "type": "bool"
        }
      ]
    },
    {
      "name": "current_election_state_pending_date",
      "base": ""
    },
    {
      "name": "current_election_state_registration_v0",
      "base": "",
      "fields": [
        {
          "name": "start_time",
          "type": "block_timestamp_type"
        },
        {
          "name": "election_threshold",
          "type": "uint16"
        }
      ]
    },
    {
      "name": "election_seeder",
      "base": "",
      "fields": [
        {
          "name": "current",
          "type": "checksum256"
        },
        {
          "name": "start_time",
          "type": "block_timestamp_type"
        },
        {
          "name": "end_time",
          "type": "block_timestamp_type"
        }
      ]
    },
    {
      "name": "current_election_state_seeding_v0",
      "base": "",
      "fields": [
        {
          "name": "seed",
          "type": "election_seeder"
        }
      ]
    },
    {
      "name": "election_rng",
      "base": "",
      "fields": [
        {
          "name": "buf",
          "type": "uint8[]"
        },
        {
          "name": "index",
          "type": "uint8"
        }
      ]
    },
    {
      "name": "current_election_state_init_voters_v0",
      "base": "",
      "fields": [
        {
          "name": "next_member_idx",
          "type": "uint16"
        },
        {
          "name": "rng",
          "type": "election_rng"
        },
        {
          "name": "last_processed",
          "type": "name"
        },
        {
          "name": "next_report_index",
          "type": "uint16"
        }
      ]
    },
    {
      "name": "election_round_config",
      "base": "",
      "fields": [
        {
          "name": "num_participants",
          "type": "uint16"
        },
        {
          "name": "num_groups",
          "type": "uint16"
        }
      ]
    },
    {
      "name": "current_election_state_active",
      "base": "",
      "fields": [
        {
          "name": "round",
          "type": "uint8"
        },
        {
          "name": "config",
          "type": "election_round_config"
        },
        {
          "name": "saved_seed",
          "type": "checksum256"
        },
        {
          "name": "round_end",
          "type": "block_timestamp_type"
        }
      ]
    },
    {
      "name": "current_election_state_post_round",
      "base": "",
      "fields": [
        {
          "name": "rng",
          "type": "election_rng"
        },
        {
          "name": "prev_round",
          "type": "uint8"
        },
        {
          "name": "prev_config",
          "type": "election_round_config"
        },
        {
          "name": "next_input_index",
          "type": "uint16"
        },
        {
          "name": "next_output_index",
          "type": "uint16"
        },
        {
          "name": "next_report_index",
          "type": "uint16"
        }
      ]
    },
    {
      "name": "current_election_state_final",
      "base": "",
      "fields": [
        {
          "name": "seed",
          "type": "election_seeder"
        }
      ]
    },
    {
      "name": "current_election_state_registration_v1",
      "base": "",
      "fields": [
        {
          "name": "start_time",
          "type": "block_timestamp_type"
        },
        {
          "name": "election_threshold",
          "type": "uint16"
        },
        {
          "name": "election_schedule_version",
          "type": "uint8"
        }
      ]
    },
    {
      "name": "current_election_state_seeding_v1",
      "base": "",
      "fields": [
        {
          "name": "seed",
          "type": "election_seeder"
        },
        {
          "name": "election_schedule_version",
          "type": "uint8"
        }
      ]
    },
    {
      "name": "current_election_state_init_voters_v1",
      "base": "",
      "fields": [
        {
          "name": "next_member_idx",
          "type": "uint16"
        },
        {
          "name": "rng",
          "type": "election_rng"
        },
        {
          "name": "last_processed",
          "type": "name"
        },
        {
          "name": "next_report_index",
          "type": "uint16"
        },
        {
          "name": "election_schedule_version",
          "type": "uint8"
        }
      ]
    },
    {
      "name": "election_state_v0",
      "base": "",
      "fields": [
        {
          "name": "lead_representative",
          "type": "name"
        },
        {
          "name": "board",
          "type": "name[]"
        },
        {
          "name": "last_election_time",
          "type": "block_timestamp_type"
        }
      ]
    },
    {
      "name": "encrypted_data_v0",
      "base": "",
      "fields": [
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "keys",
          "type": "encrypted_key[]"
        },
        {
          "name": "data",
          "type": "bytes"
        }
      ]
    },
    {
      "name": "global_data_v0",
      "base": "",
      "fields": [
        {
          "name": "community",
          "type": "string"
        },
        {
          "name": "minimum_donation",
          "type": "asset"
        },
        {
          "name": "auction_starting_bid",
          "type": "asset"
        },
        {
          "name": "auction_duration",
          "type": "uint32"
        },
        {
          "name": "stage",
          "type": "uint8"
        }
      ]
    },
    {
      "name": "global_data_v1",
      "base": "",
      "fields": [
        {
          "name": "community",
          "type": "string"
        },
        {
          "name": "minimum_donation",
          "type": "asset"
        },
        {
          "name": "auction_starting_bid",
          "type": "asset"
        },
        {
          "name": "auction_duration",
          "type": "uint32"
        },
        {
          "name": "stage",
          "type": "uint8"
        },
        {
          "name": "election_start_time",
          "type": "uint32"
        },
        {
          "name": "election_round_time_sec",
          "type": "uint32"
        }
      ]
    },
    {
      "name": "induction_v0",
      "base": "",
      "fields": [
        {
          "name": "id",
          "type": "uint64"
        },
        {
          "name": "inviter",
          "type": "name"
        },
        {
          "name": "invitee",
          "type": "name"
        },
        {
          "name": "endorsements",
          "type": "uint32"
        },
        {
          "name": "created_at",
          "type": "block_timestamp_type"
        },
        {
          "name": "video",
          "type": "string"
        },
        {
          "name": "new_member_profile",
          "type": "new_member_profile"
        }
      ]
    },
    {
      "name": "member_v0",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "name",
          "type": "string"
        },
        {
          "name": "status",
          "type": "uint8"
        },
        {
          "name": "nft_template_id",
          "type": "uint64"
        }
      ]
    },
    {
      "name": "member_v1",
      "base": "",
      "fields": [
        {
          "name": "account",
          "type": "name"
        },
        {
          "name": "name",
          "type": "string"
        },
        {
          "name": "status",
          "type": "uint8"
        },
        {
          "name": "nft_template_id",
          "type": "uint64"
        },
        {
          "name": "election_participation_status",
          "type": "uint8"
        },
        {
          "name": "election_rank",
          "type": "uint8"
        },
        {
          "name": "representative",
          "type": "name"
        },
        {
          "name": "encryption_key",
          "type": "public_key?"
        }
      ]
    },
    {
      "name": "member_stats_v0",
      "base": "",
      "fields": [
        {
          "name": "active_members",
          "type": "uint16"
        },
        {
          "name": "pending_members",
          "type": "uint16"
        },
        {
          "name": "completed_waiting_inductions",
          "type": "uint16"
        }
      ]
    },
    {
      "name": "member_stats_v1",
      "base": "",
      "fields": [
        {
          "name": "active_members",
          "type": "uint16"
        },
        {
          "name": "pending_members",
          "type": "uint16"
        },
        {
          "name": "completed_waiting_inductions",
          "type": "uint16"
        },
        {
          "name": "ranks",
          "type": "uint16[]"
        }
      ]
    },
    {
      "name": "migrate_auction_v0",
      "base": "",
      "fields": [
        {
          "name": "last_auction_id",
          "type": "uint64"
        }
      ]
    },
    {
      "name": "migrate_account_v0",
      "base": "",
      "fields": [
        {
          "name": "last_visited",
          "type": "name"
        },
        {
          "name": "user_total",
          "type": "asset"
        }
      ]
    },
    {
      "name": "no_migration<0>",
      "base": ""
    },
    {
      "name": "migrate_member_v0",
      "base": "",
      "fields": [
        {
          "name": "next_primary_key",
          "type": "uint64"
        }
      ]
    },
    {
      "name": "no_migration<1>",
      "base": ""
    },
    {
      "name": "no_migration<2>",
      "base": ""
    },
    {
      "name": "pool_v0",
      "base": "",
      "fields": [
        {
          "name": "name",
          "type": "name"
        },
        {
          "name": "monthly_distribution_pct",
          "type": "uint8"
        }
      ]
    },
    {
      "name": "session_v0",
      "base": "",
      "fields": [
        {
          "name": "key",
          "type": "public_key"
        },
        {
          "name": "expiration",
          "type": "block_timestamp_type"
        },
        {
          "name": "description",
          "type": "string"
        },
        {
          "name": "sequences",
          "type": "varuint32[]"
        }
      ]
    },
    {
      "name": "session_container_v0",
      "base": "",
      "fields": [
        {
          "name": "eden_account",
          "type": "name"
        },
        {
          "name": "earliest_expiration",
          "type": "block_timestamp_type"
        },
        {
          "name": "sessions",
          "type": "session_v0[]"
        }
      ]
    },
    {
      "name": "vote",
      "base": "",
      "fields": [
        {
          "name": "member",
          "type": "name"
        },
        {
          "name": "round",
          "type": "uint8"
        },
        {
          "name": "index",
          "type": "uint16"
        },
        {
          "name": "candidate",
          "type": "name"
        }
      ]
    }
  ],
  "actions": [
    {
      "name": "withdraw",
      "type": "withdraw",
      "ricardian_contract": "---\nspec_version: \"0.2.0\"\ntitle: Withdraw from my Eden balance\nsummary: Withdraw {{nowrap quantity}} from my Eden balance\nicon: https://ipfs.io/ipfs/QmToeuuNcTXgZPhGLShi9E18qFyQfr92b8fmjWS3roJwq5#aae9c37e262c08f73151a6d415df37d4317de36d76aabfaa1a6249cfdeaffeb2\n---\n\nI am withdrawing {{quantity}} from my available Eden balance."
    },
    {
      "name": "donate",
      "type": "donate",
      "ricardian_contract": ""
    },
    {
      "name": "fundtransfer",
      "type": "fundtransfer",
      "ricardian_contract": ""
    },
    {
      "name": "usertransfer",
      "type": "usertransfer",
      "ricardian_contract": ""
    },
    {
      "name": "genesis",
      "type": "genesis",
      "ricardian_contract": "---\nspec_version: \"0.2.0\"\ntitle: Start an Eden Community\nsummary: Found a new community with Genesis members\nicon: https://ipfs.io/ipfs/QmToeuuNcTXgZPhGLShi9E18qFyQfr92b8fmjWS3roJwq5#aae9c37e262c08f73151a6d415df37d4317de36d76aabfaa1a6249cfdeaffeb2\n---\n\nI hereby start a new Eden community named {{community}} with the following founding Peace Treaty and Bylaws:\n\n## Peace Treaty\n{{$clauses.peacetreaty}}\n\n## Bylaws\n{{$clauses.bylaws}}"
    },
    {
      "name": "setmindonfee",
      "type": "setmindonfee",
      "ricardian_contract": ""
    },
    {
      "name": "addtogenesis",
      "type": "addtogenesis",
      "ricardian_contract": ""
    },
    {
      "name": "gensetexpire",
      "type": "gensetexpire",
      "ricardian_contract": ""
    },
    {
      "name": "clearall",
      "type": "clearall",
      "ricardian_contract": "---\nspec_version: \"0.2.0\"\ntitle: Clear Eden Community\nsummary: WARNING - DELETING COMMUNITY RECORDS\nicon: https://ipfs.io/ipfs/QmToeuuNcTXgZPhGLShi9E18qFyQfr92b8fmjWS3roJwq5#aae9c37e262c08f73151a6d415df37d4317de36d76aabfaa1a6249cfdeaffeb2\n---\n\nI hereby clear and remove all tables relating to this Eden community. I affirm that I am authorized to do so by the aforementioned community. I understand that this will remove all community members, and destroy community information stored in chain state. Member information will, however, remain in the blockchain history."
    },
    {
      "name": "inductinit",
      "type": "inductinit",
      "ricardian_contract": "---\nspec_version: \"0.2.0\"\ntitle: Extend Eden Invitation\nsummary: Invite someone into the community\nicon: https://ipfs.io/ipfs/QmToeuuNcTXgZPhGLShi9E18qFyQfr92b8fmjWS3roJwq5#aae9c37e262c08f73151a6d415df37d4317de36d76aabfaa1a6249cfdeaffeb2\n---\n\nAs an active member of this Eden community, I extend an invitation to {{invitee}} to join the Eden community, pending {{invitee}}'s completion of the induction process as witnessed by the following other currently-active community members: {{witnesses.[0]}} and {{witnesses.[1]}}."
    },
    {
      "name": "inductmeetin",
      "type": "inductmeetin",
      "ricardian_contract": ""
    },
    {
      "name": "inductprofil",
      "type": "inductprofil",
      "ricardian_contract": "---\nspec_version: \"0.2.0\"\ntitle: Create My Eden Profile\nsummary: Affirm profile, Peace Treaty and Bylaws\nicon: https://ipfs.io/ipfs/QmToeuuNcTXgZPhGLShi9E18qFyQfr92b8fmjWS3roJwq5#aae9c37e262c08f73151a6d415df37d4317de36d76aabfaa1a6249cfdeaffeb2\n---\n\nI, {{new_member_profile.name}}, certify that, to the best of my knowledge, the profile information I am submitting herein is accurate and is my own. I affirm and agree to abide by this Eden community's Peace Treaty and Bylaws:\n\n## Peace Treaty\n{{$clauses.peacetreaty}}\n\n## Bylaws\n{{$clauses.bylaws}}"
    },
    {
      "name": "inductvideo",
      "type": "inductvideo",
      "ricardian_contract": "---\nspec_version: \"0.2.0\"\ntitle: Add Induction Ceremony to the Record\nsummary: Add video recording of invitee's induction ceremony to the record\nicon: https://ipfs.io/ipfs/QmToeuuNcTXgZPhGLShi9E18qFyQfr92b8fmjWS3roJwq5#aae9c37e262c08f73151a6d415df37d4317de36d76aabfaa1a6249cfdeaffeb2\n---\n\nI witnessed the Eden induction ceremony for the invitee in Induction #{{id}} and hereby attach the IPFS CID of said video recording hereto."
    },
    {
      "name": "inductendors",
      "type": "inductendors",
      "ricardian_contract": "---\nspec_version: \"0.2.0\"\ntitle: Endorse Prospective Eden Member\nsummary: Endorsement of invitee for induction into Eden community\nicon: https://ipfs.io/ipfs/QmToeuuNcTXgZPhGLShi9E18qFyQfr92b8fmjWS3roJwq5#aae9c37e262c08f73151a6d415df37d4317de36d76aabfaa1a6249cfdeaffeb2\n---\n\nI witnessed the Eden induction ceremony for the invitee in Induction #{{id}}. I believe they understand the Peace Treaty and will abide by it. I have carefully reviewed the prospective member's profile information, including their name, profile statement, social links, and induction ceremony video recording, and I affirm their accuracy to the best of my knowledge. I hereby endorse this prospective Eden member for induction into this Eden community in accordance with the Peace Treaty and Bylaws:\n\n## Peace Treaty\n{{$clauses.peacetreaty}}\n\n## Bylaws\n{{$clauses.bylaws}}"
    },
    {
      "name": "setencpubkey",
      "type": "setencpubkey",
      "ricardian_contract": ""
    },
    {
      "name": "electsettime",
      "type": "electsettime",
      "ricardian_contract": ""
    },
    {
      "name": "electconfig",
      "type": "electconfig",
      "ricardian_contract": ""
    },
    {
      "name": "electopt",
      "type": "electopt",
      "ricardian_contract": ""
    },
    {
      "name": "electseed",
      "type": "electseed",
      "ricardian_contract": ""
    },
    {
      "name": "electmeeting",
      "type": "electmeeting",
      "ricardian_contract": ""
    },
    {
      "name": "electvote",
      "type": "electvote",
      "ricardian_contract": ""
    },
    {
      "name": "electvideo",
      "type": "electvideo",
      "ricardian_contract": ""
    },
    {
      "name": "electprocess",
      "type": "electprocess",
      "ricardian_contract": ""
    },
    {
      "name": "bylawspropose",
      "type": "bylawspropose",
      "ricardian_contract": ""
    },
    {
      "name": "bylawsapprove",
      "type": "bylawsapprove",
      "ricardian_contract": ""
    },
    {
      "name": "bylawsratify",
      "type": "bylawsratify",
      "ricardian_contract": ""
    },
    {
      "name": "distribute",
      "type": "distribute",
      "ricardian_contract": ""
    },
    {
      "name": "givesbt",
      "type": "givesbt",
      "ricardian_contract": ""
    },
    {
      "name": "setdistpct",
      "type": "setdistpct",
      "ricardian_contract": ""
    },
    {
      "name": "inductdonate",
      "type": "inductdonate",
      "ricardian_contract": "---\nspec_version: \"0.2.0\"\ntitle: Donate to the Eden Community\nsummary: Submit {{nowrap quantity}} donation and activate your membership\nicon: https://ipfs.io/ipfs/QmToeuuNcTXgZPhGLShi9E18qFyQfr92b8fmjWS3roJwq5#aae9c37e262c08f73151a6d415df37d4317de36d76aabfaa1a6249cfdeaffeb2\n---\n\nI have carefully reviewed the profile information I submitted, including my name, profile statement, and social links, and I affirm their accuracy to the best of my knowledge. I hereby donate {{quantity}} to this Eden community. I affirm that I have read and understand the Eden Peace Treaty. I agree to abide by the Peace Treaty.\n\n## Peace Treaty\n{{$clauses.peacetreaty}}\n\n## Bylaws\n{{$clauses.bylaws}}"
    },
    {
      "name": "inductcancel",
      "type": "inductcancel",
      "ricardian_contract": "---\nspec_version: \"0.2.0\"\ntitle: Cancel Induction\nsummary: Cancel Induction {{nowrap id}}\nicon: https://ipfs.io/ipfs/QmToeuuNcTXgZPhGLShi9E18qFyQfr92b8fmjWS3roJwq5#aae9c37e262c08f73151a6d415df37d4317de36d76aabfaa1a6249cfdeaffeb2\n---\n\nCancel induction (pending invitation) #{{id}}. Only the inviter or a witness can cancel the pending induction. This action will delete the induction record and any related witness endorsement records stored in chain state. This information will, however, remain in the blockchain history."
    },
    {
      "name": "inducted",
      "type": "inducted",
      "ricardian_contract": "---\nspec_version: \"0.2.0\"\ntitle: Inducted (Inline Action)\nsummary: Internal inline action\nicon: https://ipfs.io/ipfs/QmToeuuNcTXgZPhGLShi9E18qFyQfr92b8fmjWS3roJwq5#aae9c37e262c08f73151a6d415df37d4317de36d76aabfaa1a6249cfdeaffeb2\n---\n\nThis action is not intended to be called directly. It is an inline action called at the end of eden::inductdonate that activates the member and cleans up induction tables."
    },
    {
      "name": "resign",
      "type": "resign",
      "ricardian_contract": ""
    },
    {
      "name": "removemember",
      "type": "removemember",
      "ricardian_contract": ""
    },
    {
      "name": "rename",
      "type": "rename",
      "ricardian_contract": ""
    },
    {
      "name": "gc",
      "type": "gc",
      "ricardian_contract": "---\nspec_version: \"0.2.0\"\ntitle: Garbage Collect\nsummary: Clean up expired or moot invitations and endorsements\nicon: https://ipfs.io/ipfs/QmToeuuNcTXgZPhGLShi9E18qFyQfr92b8fmjWS3roJwq5#aae9c37e262c08f73151a6d415df37d4317de36d76aabfaa1a6249cfdeaffeb2\n---\n\nRemove expired induction records, moot duplicate induction records, and related endorsement records. This is a safe-to-call, housekeeping action."
    },
    {
      "name": "migrate",
      "type": "migrate",
      "ricardian_contract": ""
    },
    {
      "name": "unmigrate",
      "type": "unmigrate",
      "ricardian_contract": ""
    }
  ],
  "tables": [
    {
      "name": "account",
      "index_type": "i64",
      "type": "variant<account_v0>"
    },
    {
      "name": "auction",
      "index_type": "i64",
      "type": "variant<auction_v0>"
    },
    {
      "name": "bylaws",
      "index_type": "i64",
      "type": "variant<bylaws_v0>"
    },
    {
      "name": "distaccount",
      "index_type": "i64",
      "type": "variant<distribution_account_v0>"
    },
    {
      "name": "distribution",
      "index_type": "i64",
      "type": "variant<next_distribution,election_distribution,current_distribution>"
    },
    {
      "name": "badge",
      "index_type": "i64",
      "type": "variant<badge_v0>"
    },
    {
      "name": "endorsement",
      "index_type": "i64",
      "type": "variant<endorsement_v0>"
    },
    {
      "name": "elect.curr",
      "index_type": "i64",
      "type": "variant<current_election_state_pending_date,current_election_state_registration_v0,current_election_state_seeding_v0,current_election_state_init_voters_v0,current_election_state_active,current_election_state_post_round,current_election_state_final,current_election_state_registration_v1,current_election_state_seeding_v1,current_election_state_init_voters_v1>"
    },
    {
      "name": "elect.state",
      "index_type": "i64",
      "type": "variant<election_state_v0>"
    },
    {
      "name": "encrypted",
      "index_type": "i64",
      "type": "variant<encrypted_data_v0>"
    },
    {
      "name": "global",
      "index_type": "i64",
      "type": "variant<global_data_v0,global_data_v1>"
    },
    {
      "name": "induction",
      "index_type": "i64",
      "type": "variant<induction_v0>"
    },
    {
      "name": "member",
      "index_type": "i64",
      "type": "variant<member_v0,member_v1>"
    },
    {
      "name": "memberstats",
      "index_type": "i64",
      "type": "variant<member_stats_v0,member_stats_v1>"
    },
    {
      "name": "migration",
      "index_type": "i64",
      "type": "variant<migrate_auction_v0,migrate_account_v0,no_migration<0>,migrate_member_v0,no_migration<1>,no_migration<2>>"
    },
    {
      "name": "pools",
      "index_type": "i64",
      "type": "variant<pool_v0>"
    },
    {
      "name": "sessions",
      "index_type": "i64",
      "type": "variant<session_container_v0>"
    },
    {
      "name": "votes",
      "index_type": "i64",
      "type": "vote"
    }
  ],
  "ricardian_clauses": [
    {
      "id": "peacetreaty",
      "body": "I. The size of an independent Eden community shall not exceed 10,000 members.\nII. Leaders shall be elected by the following process:\nII.a. Members are randomly organized into roughly equally-sized groups of 12 or fewer, where total number of groups = population / average group size.\nII.b. Each group must select a representative from their group with greater than ⅔ approval.\nII.c. The process shall then repeat, fractally, by randomly grouping the representatives approved in the previous round of elections, until a single lead representative is chosen.\nIII. Elections shall occur at least annually or may be triggered by a petition of 10% of the membership or according to the bylaws.\nIV. The community may adopt bylaws, which contain all rules, processes, regulations that are binding on anyone who wishes to remain a member.\nIV.a. The level of representation that voted in the lead representative shall be known as the Board.\nIV.b. The active and a proposed set of bylaws are indivisible, single documents, ie. the former Board cannot provide the next board numerous options.\nIV.c. Bylaws may only be ratified if they were proposed at least 3 months before the last election.\nIV.d. The lead representative can act inside the existing bylaws.\nIV.e. The board can propose a new set of bylaws. The vote to approve a proposal as well as ratify a proposed set of bylaws shall be approved by ⅔+1 vote, which shall include the lead representative.\nV. Bylaws may not override, change, eliminate, or extend the Peace Treaty.\nVI. Members must be invited according to community bylaws and can be removed according to community bylaws.\nVII. Membership is voluntary. Members may leave at any time by giving notice.\nVIII. The Peace Treaty may be amended by a ⅔+1 vote of all members."
    },
    {
      "id": "bylaws",
      "body": "I. The initial Board shall consist of the people on the Genesis Call and shall have the power to propose and ratify bylaws until such time as the first election occurs.\nII. Genesis members are to only participate in up to 15 induction meetings until membership reaches 100 members."
    }
  ],
  "variants": [
    {
      "name": "attribute_value",
      "types": [
        "int8",
        "int16",
        "int32",
        "int64",
        "uint8",
        "uint16",
        "uint32",
        "uint64",
        "float32",
        "float64",
        "string",
        "INT8_VEC",
        "INT16_VEC",
        "INT32_VEC",
        "INT64_VEC",
        "UINT8_VEC",
        "UINT16_VEC",
        "UINT32_VEC",
        "UINT64_VEC",
        "FLOAT_VEC",
        "DOUBLE_VEC",
        "STRING_VEC"
      ]
    },
    {
      "name": "run_auth",
      "types": [
        "no_auth",
        "account_auth",
        "signature_auth"
      ]
    },
    {
      "name": "verb",
      "types": [
        "unsupported_verb",
        "inductmeetin",
        "inductprofil",
        "inductvideo",
        "inductendors",
        "electopt",
        "electmeeting",
        "electvote",
        "electvideo",
        "inductcancel",
        "inductinit"
      ]
    },
    {
      "name": "variant<account_v0>",
      "types": [
        "account_v0"
      ]
    },
    {
      "name": "variant<auction_v0>",
      "types": [
        "auction_v0"
      ]
    },
    {
      "name": "variant<bylaws_v0>",
      "types": [
        "bylaws_v0"
      ]
    },
    {
      "name": "variant<distribution_account_v0>",
      "types": [
        "distribution_account_v0"
      ]
    },
    {
      "name": "variant<next_distribution,election_distribution,current_distribution>",
      "types": [
        "next_distribution",
        "election_distribution",
        "current_distribution"
      ]
    },
    {
      "name": "variant<badge_v0>",
      "types": [
        "badge_v0"
      ]
    },
    {
      "name": "variant<endorsement_v0>",
      "types": [
        "endorsement_v0"
      ]
    },
    {
      "name": "variant<current_election_state_pending_date,current_election_state_registration_v0,current_election_state_seeding_v0,current_election_state_init_voters_v0,current_election_state_active,current_election_state_post_round,current_election_state_final,current_election_state_registration_v1,current_election_state_seeding_v1,current_election_state_init_voters_v1>",
      "types": [
        "current_election_state_pending_date",
        "current_election_state_registration_v0",
        "current_election_state_seeding_v0",
        "current_election_state_init_voters_v0",
        "current_election_state_active",
        "current_election_state_post_round",
        "current_election_state_final",
        "current_election_state_registration_v1",
        "current_election_state_seeding_v1",
        "current_election_state_init_voters_v1"
      ]
    },
    {
      "name": "variant<election_state_v0>",
      "types": [
        "election_state_v0"
      ]
    },
    {
      "name": "variant<encrypted_data_v0>",
      "types": [
        "encrypted_data_v0"
      ]
    },
    {
      "name": "variant<global_data_v0,global_data_v1>",
      "types": [
        "global_data_v0",
        "global_data_v1"
      ]
    },
    {
      "name": "variant<induction_v0>",
      "types": [
        "induction_v0"
      ]
    },
    {
      "name": "variant<member_v0,member_v1>",
      "types": [
        "member_v0",
        "member_v1"
      ]
    },
    {
      "name": "variant<member_stats_v0,member_stats_v1>",
      "types": [
        "member_stats_v0",
        "member_stats_v1"
      ]
    },
    {
      "name": "variant<migrate_auction_v0,migrate_account_v0,no_migration<0>,migrate_member_v0,no_migration<1>,no_migration<2>>",
      "types": [
        "migrate_auction_v0",
        "migrate_account_v0",
        "no_migration<0>",
        "migrate_member_v0",
        "no_migration<1>",
        "no_migration<2>"
      ]
    },
    {
      "name": "variant<pool_v0>",
      "types": [
        "pool_v0"
      ]
    },
    {
      "name": "variant<session_container_v0>",
      "types": [
        "session_container_v0"
      ]
    }
  ]
}
`
}
