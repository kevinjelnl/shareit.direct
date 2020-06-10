FROM scratch
ADD main /
COPY views /views
CMD ["/main"]
