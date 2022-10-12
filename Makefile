MSD_PROTO_DIR=.
# change to your local path
# https://github.com/msd-rs/msd-proto
MSD_PROTO_REPO=${HOME}/repo/msd-rs/msd-proto
# change to current pkg path
MSD_GO_PKG=github.com/msd-rs/msd

pb:
	cp ${MSD_PROTO_REPO}/*.proto .
	protoc -I ${MSD_PROTO_DIR} \
	--go_out=.  \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go_opt=Mdataframe.proto=${MSD_GO_PKG} \
	--go_opt=Mschema.proto=${MSD_GO_PKG} \
	--go_opt=Mmsd.proto=${MSD_GO_PKG} \
	--go-grpc_opt=Mdataframe.proto=${MSD_GO_PKG} \
	--go-grpc_opt=Mschema.proto=${MSD_GO_PKG} \
	--go-grpc_opt=Mmsd.proto=${MSD_GO_PKG} \
	--go-grpc_opt=paths=source_relative \
	${MSD_PROTO_DIR}/*.proto
