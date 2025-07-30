FROM golang:1.24

WORKDIR /app

COPY . .
RUN go mod tidy

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    latexmk \
    texlive-xetex \
    texlive-latex-recommended \
    texlive-latex-extra \
    texlive-fonts-recommended \
    texlive-fonts-extra \
    texlive-lang-all \
    perl \
    make \
    && apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN ["mkdir", "/storage"]
EXPOSE 3000
# RUN air init
CMD ["go", "run", "cmd/main.go"]
