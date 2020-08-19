## 1.4.0 (Aug 19, 2020)

ENHANCEMENTS:

* Add mutex to protect against concurrent operations (GH #36)
* Build on Go 1.13

BUG FIXES:

* Update gopass dependencies


## 1.3.0 (July 24, 2020)

ENHANCEMENTS:

* Set password resources values as sensitive
* Return ResourceRead errors
* Fix build with Go 1.13
* Use new Terraform plugin SDK

## 1.2.1 (June 05, 2019)

ENHANCEMENTS:

* Use Terraform v0.12.0

## 1.2.0 (May 21, 2019)

IMPROVEMENTS:

* Update Gopass dependencies
* Use Terraform v0.12-beta's API
* Improve build config, add CI

## 1.1.1 (Sep 21, 2018)

IMPROVEMENTS:

* Support for single-line password secret

BUG FIXES:

* provider: return an error if the store is not initialized

## 1.1.0 (Jun 25, 2018)

IMPROVEMENTS:

* Support newer versions of gopass

FEATURES:

* Expose entire secret contents in `.full` data source attribute

## 1.0.1 (Mar 5, 2018)

BUG FIXES:

* datasource/passwordReturn errors when setter fails

## 1.0.0 (Oct 4, 2017)

IMPROVEMENTS:

* Port to gopass

## 0.1.4 (Jul 4, 2017)

FEATURES:

* provider: refresh password store by default

## 0.1.3 (Jul 4, 2017)

FEATURES:

* resource/password: new resource
* add tests

## 0.1.2 (Feb 15, 2017)

BUG FIXES:

* datasource/password: fix attribute name

## 0.1.1 (Feb 9, 2017)

BUG FIXES:

* datasource/password: don't fail if unmarshaling fails

## 0.1.0 (Jan 16, 2017)

* Initial release
