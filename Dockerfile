# CGO_ENABLED=0 go build main.go

FROM gcr.io/distroless/static

USER nonroot:nonroot
 
# copy compiled app
COPY --chown=nonroot:nonroot /main /main
 
# run binary; use vector form
ENTRYPOINT ["/main"]