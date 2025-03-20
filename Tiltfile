# Load the restart_process extension
load('ext://restart_process', 'docker_build_with_restart')

### K8s Config ###

# Uncomment to use secrets
# k8s_yaml('./infra/development/k8s/secrets.yaml')

k8s_yaml('./infra/development/k8s/app-config.yaml')

### End of K8s Config ###
### API Gateway ###

gateway_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/api-gateway ./services/api-gateway'
if os.name == 'nt':
  gateway_compile_cmd = './infra/development/docker/api-gateway-build.bat'

local_resource(
  'api-gateway-compile',
  gateway_compile_cmd,
  deps=['./services/api-gateway', './shared'], labels="compiles")


docker_build_with_restart(
  'ride-sharing/api-gateway',
  '.',
  entrypoint=['/app/build/api-gateway'],
  dockerfile='./infra/development/docker/api-gateway.Dockerfile',
  only=[
    './build/api-gateway',
    './shared',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./shared', '/app/shared'),
  ],
)

k8s_yaml('./infra/development/k8s/api-gateway-deployment.yaml')
k8s_resource('api-gateway', port_forwards=8081,
             resource_deps=['api-gateway-compile'], labels="services")
### End of API Gateway ###
### Trip Service ###

# Uncomment once we have a trip service

#trip_compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/trip-service ./services/trip-service/cmd/main.go'
#if os.name == 'nt':
#  trip_compile_cmd = './infra/development/docker/trip-build.bat'

# local_resource(
#   'trip-service-compile',
#   trip_compile_cmd,
#   deps=['./services/trip-service', './shared'], labels="compiles")

# docker_build_with_restart(
#   'ride-sharing/trip-service',
#   '.',
#   entrypoint=['/app/build/trip-service'],
#   dockerfile='./infra/development/docker/trip-service.Dockerfile',
#   only=[
#     './build/trip-service',
#     './shared',
#   ],
#   live_update=[
#     sync('./build', '/app/build'),
#     sync('./shared', '/app/shared'),
#   ],
# )

# k8s_yaml('./infra/development/k8s/trip-service-deployment.yaml')
# k8s_resource('trip-service', resource_deps=['trip-service-compile'], labels="services")

### End of Trip Service ###
### Web Frontend ###

docker_build(
  'ride-sharing/web',
  '.',
  dockerfile='./infra/development/docker/web.Dockerfile',
)

k8s_yaml('./infra/development/k8s/web-deployment.yaml')
k8s_resource('web', port_forwards=3000, labels="frontend")

### End of Web Frontend ###