FROM alpine:3.22.0

RUN addgroup --system appuser && \
    adduser --system -G appuser appuser

USER appuser

WORKDIR /app

ARG APP
ARG TARGETOS
ARG TARGETARCH

COPY build/${APP}-${TARGETOS}-${TARGETARCH} .

ENV APP_ENV=${APP}
ENV TARGETOS_ENV=${TARGETOS}
ENV TARGETARCH_ENV=${TARGETARCH}

CMD ["sh", "-c", "./${APP_ENV}-${TARGETOS_ENV}-${TARGETARCH_ENV}"]
