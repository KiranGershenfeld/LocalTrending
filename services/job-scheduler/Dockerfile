FROM python:3.11-slim as build
RUN apt-get update && apt-get -y install gcc libpq-dev

WORKDIR /src
COPY src .
COPY requirements.txt .

RUN pip install --upgrade pip && \
    python -m pip install -r requirements.txt && \
    pip install -e common/

WORKDIR /src
CMD ["python", "main.py"]
