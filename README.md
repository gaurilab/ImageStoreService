# ImageStoreService

Implement Image Store Service in GoLang using GIN service

### Requirements:

1) [X] Create/Delete Image Album
2) [X] Create/Delete Image in an Album
3) [X] Get a single Image in an Album
4) [X] Get All Images in an Album
5) [X] WEB UI : Implement a simple web UI where you take all the inputs from user a nd display the output
6) [X] Make the service docker ready.
7) [ ] Run this application on Kubernetes
8) [ ] Write unit tests that will test all your function with at least ten different sets.
9) [ ] Write integration tests using GINKGO framework.

### Basic Design

 To design a system to store the images, highly available data store is required.  I deally  the images can be stored on dedicated file sever / cloud service (accessible by some URL). For this particular instance using mongodb for  storing the metadata and images.   Separate storage service  should be introduced to decouple/scale  the storage needs  independently.  In the implementation its just one stateless webapp which can talk to database for storing.

### APIS

| Operations | API                             | Function                                                   |
| ---------- | ------------------------------- | ---------------------------------------------------------- |
| POST       | /create-album                   | Create Album. Generates an ID.                             |
| GET        | /create-album                   | Get all Albums Ids.                                        |
| DELETE     | /delete-album/:albumID          | Delete the Album given the album id.                       |
| GET        | /albums/:albumID                | Get Album and images id in it. Pagination not implemented. |
| POST       | /upload/:albumID                | Upload Image in Album.                                     |
| GET        | /delete-image/:albumID/:imageID | Delete image in a album.                                   |
| GET        | /images/:id                     | Gets the image.                                            |

### Tests

```docker
 docker run   -p 27017:27017  my-mongodb
```

```cpp
ImageStoreService % go run main.go

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] Loaded HTML Templates (5):
	-
	- album.html
	- create-album.html
	- index.html
	- upload.html

[GIN-debug] GET    /                         --> github.com/gaurilab/ImageStoreService/api/handlers.RenderHomePage (3 handlers)
[GIN-debug] GET    /create-album             --> github.com/gaurilab/ImageStoreService/api/handlers.RenderCreateAlbumPage (3 handlers)
[GIN-debug] POST   /create-album             --> github.com/gaurilab/ImageStoreService/api/handlers.CreateAlbum (3 handlers)
[GIN-debug] GET    /delete-album/:albumID    --> github.com/gaurilab/ImageStoreService/api/handlers.DeleteAlbum (3 handlers)
[GIN-debug] GET    /albums/:albumID          --> github.com/gaurilab/ImageStoreService/api/handlers.GetAlbum (3 handlers)
[GIN-debug] GET    /upload/:albumID          --> github.com/gaurilab/ImageStoreService/api/handlers.RenderUploadImagePage (3 handlers)
[GIN-debug] POST   /upload/:albumID          --> github.com/gaurilab/ImageStoreService/api/handlers.UploadImage (3 handlers)
[GIN-debug] GET    /delete-image/:albumID/:imageID --> github.com/gaurilab/ImageStoreService/api/handlers.DeleteImage (3 handlers)
[GIN-debug] GET    /images/:id               --> github.com/gaurilab/ImageStoreService/api/handlers.RenderImage (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :9080
```

### Building and deploying  Cluster in Kubernetes


In progress 


```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: image-db-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: image-db
  template:
    metadata:
      labels:
        app: image-db
    spec:
      containers:
        - name: image-db
          image: mongo:latest  # Specify the desired MongoDB image and version
          ports:
            - containerPort: 27017
              name: "mongodb"
          volumeMounts:
            - mountPath: "/etc/mongodb"
              name: image-data-storage

      volumes:
        - name: image-data-storage
          persistentVolumeClaim:
            claimName: image-data-persisent-volume-claim
# db-persistent-volume-claim.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: image-data-persisent-volume-claim
spec:
  volumeName: image-data-persisent-volume
  storageClassName: manual
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
# db-persistent-volume.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: image-data-persisent-volume
  labels:
    type: local
spec:
  claimRef:
    namespace: default
    name: image-data-persisent-volume-claim
  storageClassName: manual
  capacity:
    storage: 5Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/Users/gauribehera/workspace/"
# db-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: image-db-service
spec:
  type: NodePort
  selector:
    app: image-db
  ports:
    - name: "mongodb"
      protocol: TCP
      port: 27017
      targetPort: 27017
      nodePort: 32017
```



### Challenges


Unable to expose the port  from the Kubernetes Cluster with the service.
