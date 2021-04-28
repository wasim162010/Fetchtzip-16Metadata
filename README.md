# FetchURL

#This is an application which fetches metadata by hitting urls 

# To run an application, you can pass url as a command line parameter. For Ex,
#go run QueryUrl.go "https://storage.googleapis.com/tzip-16/emoji-in-metadata.json"
#go run QueryUrl.go "sha256://0x7e99ecf3a4490e3044ccdf319898d77380a2fc20aae36b6e40327d678399d17b/https:%2F%2Fstorage.googleapis.com%2Ftzip-16%2Ftaco-shop-metadata.json"


# To run test cases, you can pass the url while the go test command.It uses "flag" package to enable parameter(url in this case) as a command line param while executing the test case.
#for ex:

#go test -run TestQueryHttpUrl -httpURL "https://storage.googleapis.com/tzip-16/emoji-in-metadata.json"

#go test -run TestQuerySHA256Url -sha256URL "sha256://0x7e99ecf3a4490e3044ccdf319898d77380a2fc20aae36b6e40327d678399d17b/https:%2F%2Fstorage.googleapis.com%2Ftzip-16%2Ftaco-shop-metadata.json"

