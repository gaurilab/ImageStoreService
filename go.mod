module github.com/gaurilab/ImageStoreService

go 1.17

require (
    github.com/gin-gonic/gin v1.7.4
    github.com/onsi/ginkgo v1.16.1
    github.com/onsi/gomega v1.16.0
)

replace github.com/gaurilab/ImageStoreService/api/handlers => ./api/handlers 
replace github.com/gaurilab/ImageStoreService/api/model => ./api/model
