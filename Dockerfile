FROM scratch
COPY e-factura /e-factura
COPY templates /templates
COPY unitati.xml /unitati.xml
COPY tmp /tmp
COPY .env /.env
COPY zoneinfo.zip /
ENV TZ=Europe/Bucharest
ENV ZONEINFO=/zoneinfo.zip
ENTRYPOINT [ "/e-factura" ]