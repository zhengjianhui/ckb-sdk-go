# [v0.43.2](https://github.com/nervosnetwork/ckb-sdk-java/compare/v0.43.2-rc.1...v0.43.2) (2021-09-10)

### Features

* Add `build_smart_transfer_transaction`
  api ([9ad61d3](https://github.com/nervosnetwork/ckb-sdk-go/pull/59/commits/9ad61d3f69422c5fc58b15830942bec333294eb9))

### BreakingChanges

* The uniform unit is
  shannon ([c1492d4](https://github.com/zhengjianhui/ckb-sdk-go/commit/c1492d4bd2450b39cf3daeb1e507119ce72d0036))

* Adjusting the `build_asset_account_creation_transaction`
  interface ([de4be38](https://github.com/nervosnetwork/ckb-sdk-go/pull/54/commits/de4be388808f7f5d7cd8b810addf9320c8b36ad6))

* Adjusting the `get_generic_transaction`
  interface ([91ba601](https://github.com/nervosnetwork/ckb-sdk-go/pull/55/commits/91ba6010b8908203c97cfc4f8ef1fa85a13f6642))

* Adjusting the `get_generic_block`
  interface ([3c4ffbc](https://github.com/nervosnetwork/ckb-sdk-go/pull/57/commits/3c4ffbc372dbf72faba2d074ec40b3185d13c97c))

* Adjusting the `get_balance`
  interface ([361e7f0](https://github.com/nervosnetwork/ckb-sdk-go/pull/53/commits/361e7f0017ae725e0fbf601e4d3327ddddb13977))

# [v0.43.2-rc.1](https://github.com/nervosnetwork/ckb-sdk-go/compare/v0.43.1...v0.43.2-rc.1) (2021-08-06)

### Bug Fixes

* Fix capacity overflow on
  payment ([98aff7e](https://github.com/nervosnetwork/ckb-sdk-go/pull/51/commits/98aff7ec39e1680bd47be1e3bb623893090b795e))

# [v0.43.1](https://github.com/nervosnetwork/ckb-sdk-go/compare/v0.43.0...v0.43.1) (2021-08-01)

### Features

* Relay ckb rpc request through
  mercury ([d0e4c0a](https://github.com/nervosnetwork/ckb-sdk-go/pull/35/commits/d0e4c0a2ea2454aae52e6ff0d87d96832f18b2c1)
  and [3c7ce9b](https://github.com/nervosnetwork/ckb-sdk-go/pull/49/commits/3c7ce9b1a5cedc8aaecfd745c7d9b2f1e063f6a5))
* Add `get_fork_block`
  api ([56f578e](https://github.com/nervosnetwork/ckb-sdk-go/pull/48/commits/56f578e9cd6b9acf46642971cf9c117eaa81b6b7))
* Add `get_block_median_time`
  api ([464cf71](https://github.com/nervosnetwork/ckb-sdk-go/pull/47/commits/464cf71230738f588d46b624db7597539bfdb4d9))
* Add `clear_tx_pool`
  api ([a4aea75](https://github.com/nervosnetwork/ckb-sdk-go/pull/46/commits/a4aea75f46c62a0f7a053f709d60c37e694c5d81))
* Add `raw_tx_pool`
  api ([0ba6fa5](https://github.com/nervosnetwork/ckb-sdk-go/pull/45/commits/0ba6fa563c468b39b7ad42c9880f72de434f0845))
* Add `ping_peers`
  api ([9a8ca27](https://github.com/nervosnetwork/ckb-sdk-go/pull/44/commits/9a8ca276aff50259bd8a4595d21c7ea8d916aded))
* Add `remove_node`
  api ([95c7b05](https://github.com/nervosnetwork/ckb-sdk-go/pull/43/commits/95c7b059c06cfc5aec6c9f91a6edb839e3e9bd18))
* Add `add_node`
  api ([9e35021](https://github.com/nervosnetwork/ckb-sdk-go/pull/42/commits/9e3502161064e2eb8012becca2d87f8a55c288aa))
* Add `clear_banned_address`
  api ([9fe3589](https://github.com/nervosnetwork/ckb-sdk-go/pull/41/commits/9fe3589d3bc6479b24d98d6f4d2d4837c6c96a63))
* Add `verify_transaction_proof`
  api ([e270c42](https://github.com/nervosnetwork/ckb-sdk-go/pull/40/commits/e270c423da0e143c6d2dd2574cde7aa6c2f73aae))
* Add `set_network_active`
  api ([e704431](https://github.com/nervosnetwork/ckb-sdk-go/pull/39/commits/e704431b1ccb63f14db715fae74fd7c94638f10d))
* Add `sync_state`
  api ([c4131be](https://github.com/nervosnetwork/ckb-sdk-go/pull/38/commits/c4131bec2722f4b0764d7a7c475bee28a688f9a5))
* Add `get_transaction_proof`
  api ([18209e9](https://github.com/nervosnetwork/ckb-sdk-go/pull/37/commits/18209e9af95302d191e2656eb8b5580ff2108437))

# [v0.43.0](https://github.com/nervosnetwork/ckb-sdk-go/compare/v0.42.0...v0.43.0) (2021-07-25)

### Features

* [#30](https://github.com/nervosnetwork/ckb-sdk-go/pull/30) Support Mercury version 0.1.0-rc2
* Support CKB version 0.43.0
* [#30](https://github.com/nervosnetwork/ckb-sdk-go/pull/30) Support ckb-indexer version 0.2.1

# [v0.42.0](https://github.com/nervosnetwork/ckb-sdk-go/compare/v0.41.0...v0.42.0) (2021-06-25)

### BreakingChanges

* Remove `get_cellbase_output_capacity_details` rpc

### Bug Fixes

* Fix the locked field of the Balance struct to Locked

# [v0.41.0](https://github.com/nervosnetwork/ckb-sdk-go/compare/v0.3.0...v0.41.0) (2021-06-21)

### Features

* Implement Mercury
  SDK ([2b29782](https://github.com/nervosnetwork/ckb-sdk-go/commit/2b2978252036d264b9a1d92368b57222e75d6e34))
* SDK version follows CKB version

# [v0.3.0](https://github.com/nervosnetwork/ckb-sdk-go/compare/v0.2.1...v0.3.0) (2021-03-11)

### Bug Fixes

* wrong change target ([0706e68](https://github.com/nervosnetwork/ckb-sdk-go/commit/0706e68))
* wrong witness ([184145b](https://github.com/nervosnetwork/ckb-sdk-go/commit/184145b))

### Features

* add cheque cell args generation function ([a4d6f63](https://github.com/nervosnetwork/ckb-sdk-go/commit/a4d6f63))
* add claim cheque payment ([5579e55](https://github.com/nervosnetwork/ckb-sdk-go/commit/5579e55))
* add claim cheques unsigned tx builder ([065848d](https://github.com/nervosnetwork/ckb-sdk-go/commit/065848d))
* add GetBlockEconomicState ([0edae6a](https://github.com/nervosnetwork/ckb-sdk-go/commit/0edae6a))
* add IsChequeCell function ([8af982e](https://github.com/nervosnetwork/ckb-sdk-go/commit/8af982e))
* add IssuingChequeUnsignedTxBuilder ([76c964d](https://github.com/nervosnetwork/ckb-sdk-go/commit/76c964d))
* add OutputsCapacity method to transaction ([5702cd0](https://github.com/nervosnetwork/ckb-sdk-go/commit/5702cd0))
* add ReceiverInfo ([d99b073](https://github.com/nervosnetwork/ckb-sdk-go/commit/d99b073))
* add searchLimit const ([03398d7](https://github.com/nervosnetwork/ckb-sdk-go/commit/03398d7))
* add sign issue cheque tx method ([0ccf34c](https://github.com/nervosnetwork/ckb-sdk-go/commit/0ccf34c))
* add since generator ([69a83c4](https://github.com/nervosnetwork/ckb-sdk-go/commit/69a83c4))
* add sudt payment ([4938ff8](https://github.com/nervosnetwork/ckb-sdk-go/commit/4938ff8))
* add udt live cell collector ([73b7898](https://github.com/nervosnetwork/ckb-sdk-go/commit/73b7898))
* add UnsignedTxBuilder interface ([1c5588e](https://github.com/nervosnetwork/ckb-sdk-go/commit/1c5588e))
* add ValidateChequeAddress function ([a8621ff](https://github.com/nervosnetwork/ckb-sdk-go/commit/a8621ff))
* add withdraw cheque payment ([611b7d3](https://github.com/nervosnetwork/ckb-sdk-go/commit/611b7d3))
* hash witnesses which do not in any input group ([e3a293f](https://github.com/nervosnetwork/ckb-sdk-go/commit/e3a293f))
* implement generate unsigned issuing cheque tx ([d2d09d7](https://github.com/nervosnetwork/ckb-sdk-go/commit/d2d09d7))
* more set function on SystemScript ([129dec1](https://github.com/nervosnetwork/ckb-sdk-go/commit/129dec1))
* remove get_cells_by_lock_hash RPC and RPCs under indexer
  module ([fc920d4](https://github.com/nervosnetwork/ckb-sdk-go/commit/fc920d4))
* support custom SystemScripts ([f851553](https://github.com/nervosnetwork/ckb-sdk-go/commit/f851553))
* support filter on ckb-indexer searchKey ([aebf2f7](https://github.com/nervosnetwork/ckb-sdk-go/commit/aebf2f7))

### BREAKING CHANGES

* Remove RPCs under indexer module
* need send a SystemScripts to `GenerateTx` method manually

# [v0.2.1](https://github.com/nervosnetwork/ckb-sdk-go/compare/v0.2.0...v0.2.1) (2020-11-25)

### Features

* [#13](https://github.com/nervosnetwork/ckb-sdk-go/pull/13): expose GenerateFullPayloadAddress function

# [v0.2.0](https://github.com/nervosnetwork/ckb-bitpie-sdk/compare/v0.1.0...v0.2.0) (2020-11-25)

### Bug Fixes

* [#5](https://github.com/nervosnetwork/ckb-sdk-go/pull/5): fix nil pointer dereference on toCellWithStatus function
* [#7](https://github.com/nervosnetwork/ckb-sdk-go/pull/7): fix tx fee calculation bug

### Features

* [#6](https://github.com/nervosnetwork/ckb-sdk-go/pull/6): support ckb indexer
* [#8](https://github.com/nervosnetwork/ckb-sdk-go/pull/8): support ckb indexer on payment
* [#9](https://github.com/nervosnetwork/ckb-sdk-go/pull/9): add OccupiedCapacity function
* [#10](https://github.com/nervosnetwork/ckb-sdk-go/pull/10): support generate and parse short acp address
