FROM mcr.microsoft.com/dotnet/sdk:5.0 as build

WORKDIR /project

COPY ./services/vesselMotionService/vesselMotionService.csproj ./services/vesselMotionService/

RUN dotnet restore ./services/vesselMotionService/vesselMotionService.csproj

COPY ./services/vesselMotionService ./services/vesselMotionService

RUN dotnet publish -c Release -o /published services/vesselMotionService/vesselMotionService.csproj

FROM mcr.microsoft.com/dotnet/aspnet:5.0 as runtime

# Uncomment the line below if running with HTTPS
# ENV ASPNETCORE_URLS=https://+:443

WORKDIR /app

COPY --from=build /published .

ENTRYPOINT [ "dotnet", "vesselMotionService.dll" ]