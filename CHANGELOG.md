<a name="unreleased"></a>
## [Unreleased]

### Features
- create public model ([#71](https://github.com/SAONetwork/sao-consensus/issues/71))
- beta release ([#67](https://github.com/SAONetwork/sao-consensus/issues/67))
- add builtin did to support read public data ([#65](https://github.com/SAONetwork/sao-consensus/issues/65))

### Bug Fixes
- timeout, migrate, renew and complete  ([#66](https://github.com/SAONetwork/sao-consensus/issues/66))

### Code Refactoring
- change metadata and order to completed state when one shard in the order is completed ([#64](https://github.com/SAONetwork/sao-consensus/issues/64))


<a name="testnet-v0.1.7"></a>
## [testnet-v0.1.7](https://github.com/SAONetwork/sao-consensus/compare/v0.1.7...testnet-v0.1.7) (2023-08-28)

### Features

* store the mapping of account and key did ([#73](https://github.com/SAONetwork/sao-consensus/issues/73))  *#73* 
* list shard by sp ([#75](https://github.com/SAONetwork/sao-consensus/issues/75))  *#75* 
* create public model ([#71](https://github.com/SAONetwork/sao-consensus/issues/71))  *#71* 
* add builtin did to support read public data ([#65](https://github.com/SAONetwork/sao-consensus/issues/65))  *#65* 

### Bug Fixes

* fix shard status  ([#69](https://github.com/SAONetwork/sao-consensus/issues/69))  *#69* 
* timeout, migrate, renew and complete  ([#66](https://github.com/SAONetwork/sao-consensus/issues/66))  *#66* 

### Code Refactoring

* change metadata and order to completed state when one shard in the order is completed ([#64](https://github.com/SAONetwork/sao-consensus/issues/64))  *#64* 


<a name="v0.1.7"></a>
## [v0.1.7](https://github.com/SAONetwork/sao-consensus/compare/v0.1.6...v0.1.7) (2023-06-29)

### Features

* beta release ([#67](https://github.com/SAONetwork/sao-consensus/issues/67))  *#67* 
* add builtin did to support read public data ([#65](https://github.com/SAONetwork/sao-consensus/issues/65))  *#65* 

### Bug Fixes

* timeout, migrate, renew and complete  ([#66](https://github.com/SAONetwork/sao-consensus/issues/66))  *#66* 

### Code Refactoring

* change metadata and order to completed state when one shard in the order is completed ([#64](https://github.com/SAONetwork/sao-consensus/issues/64))  *#64* 


<a name="v0.1.6"></a>
## [v0.1.6](https://github.com/SAONetwork/sao-consensus/compare/v0.1.5...v0.1.6) (2023-06-15)

### Features

* proof of existence ([#57](https://github.com/SAONetwork/sao-consensus/issues/57))  *#57* 
* add total storage and used storage in pledge ([#58](https://github.com/SAONetwork/sao-consensus/issues/58))  *#58* 
* allow sp to declare delegation with validator for higher order priority ([#54](https://github.com/SAONetwork/sao-consensus/issues/54))  *#54* 

### Bug Fixes

* add shard pledge field to record order pledge ([#60](https://github.com/SAONetwork/sao-consensus/issues/60))  *#60* 


<a name="v0.1.5"></a>
## [v0.1.5](https://github.com/SAONetwork/sao-consensus/compare/v0.1.4...v0.1.5) (2023-05-31)

### Features

* add node params ([#53](https://github.com/SAONetwork/sao-consensus/issues/53))  *#53* 
* renew info ([#52](https://github.com/SAONetwork/sao-consensus/issues/52))  *#52* 
* pledge_debt, refactor node module funcs ([#51](https://github.com/SAONetwork/sao-consensus/issues/51))  *#51* 
* expired shard ([#49](https://github.com/SAONetwork/sao-consensus/issues/49))  *#49* 
* storage pledge baseline ([#48](https://github.com/SAONetwork/sao-consensus/issues/48))  *#48*  *#2* 

### Bug Fixes

* load metadata with data owner 
* export/import issue ([#44](https://github.com/SAONetwork/sao-consensus/issues/44))  *#44* 

### Code Refactoring

* renew ([#45](https://github.com/SAONetwork/sao-consensus/issues/45))  *#45* 
* shard record storage duration ([#46](https://github.com/SAONetwork/sao-consensus/issues/46))  *#46*  *#3* 


<a name="v0.1.4"></a>
## [v0.1.4](https://github.com/SAONetwork/sao-consensus/compare/v0.1.3...v0.1.4) (2023-04-20)

### Features

* migration for v0.1.4 ([#41](https://github.com/SAONetwork/sao-consensus/issues/41))  *#0*  *#1*  *#4*  *#41*  *#014* 
* handle timeout order ([#37](https://github.com/SAONetwork/sao-consensus/issues/37))  *#37* 

### Bug Fixes

* eth address binding proof verify ([#42](https://github.com/SAONetwork/sao-consensus/issues/42))  *#42* 
* update timeout shards ([#40](https://github.com/SAONetwork/sao-consensus/issues/40))  *#40* 
* add ignore list in random sp for migration and resize ([#33](https://github.com/SAONetwork/sao-consensus/issues/33))  *#33* 
* fill migrate shard sp and size ([#32](https://github.com/SAONetwork/sao-consensus/issues/32))  *#32* 
* add worker migrate logic ([#30](https://github.com/SAONetwork/sao-consensus/issues/30))  *#30* 
* shard pledge reward debt ([#29](https://github.com/SAONetwork/sao-consensus/issues/29))  *#29* 
* use block height to calculate reward ([#27](https://github.com/SAONetwork/sao-consensus/issues/27))  *#27* 
* remove reset_store ([#25](https://github.com/SAONetwork/sao-consensus/issues/25))  *#25* 

### Code Refactoring

* node info enhancement for UI display and tx address pool supporting ([#39](https://github.com/SAONetwork/sao-consensus/issues/39))  *#39* 
* metadata ([#38](https://github.com/SAONetwork/sao-consensus/issues/38))  *#38* 
* remove order instead of set cancel/terminate status ([#36](https://github.com/SAONetwork/sao-consensus/issues/36))  *#36* 
* data management ([#35](https://github.com/SAONetwork/sao-consensus/issues/35))  *#35* 
* uniform order status ([#31](https://github.com/SAONetwork/sao-consensus/issues/31))  *#31* 
* pool_management ([#28](https://github.com/SAONetwork/sao-consensus/issues/28))  *#28* 


<a name="v0.1.3"></a>
## [v0.1.3](https://github.com/SAONetwork/sao-consensus/compare/v0.1.2...v0.1.3) (2023-03-15)

