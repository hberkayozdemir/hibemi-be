load('ext://restart_process', 'docker_build_with_restart')
load('ext://local_output', 'local_output')

# Helper Functions

def blob_to_str(blob):
  return str(blob)[0:len(str(blob))-1]

def get_cmd_output(cmd):
  return blob_to_str(local(cmd))

def install_dev_dependencies(deps):
  for command in deps:
    out = local('command -v {} || echo "not found"'.format(command))
    if blob_to_str(out) == "not found":
      print("{} command not exist.".format(command))
      print("Installing {} package".format(command))
      local(dependencies[command])
    else:
      print("{} found!".format(command))

def start_minikube(desired_memory_capacity):
  current_memory_capacity = get_cmd_output('minikube config get memory')

  if int(current_memory_capacity) < int(desired_memory_capacity):
    local('minikube config set memory {}'.format(desired_memory_capacity))
    local('minikube delete')
    local('minikube start')
  else:
    number_of_running_words = int(get_cmd_output('minikube status | grep Running | wc -l ').lstrip(' '))
    if not number_of_running_words == 3:
      print("Starting minikube....")
      local('minikube start')

dependencies = {
  "go": "brew install go@1.15",
  "golangci-lint": "brew install golangci/tap/golangci-lint",
}

install_dev_dependencies(dependencies)

start_minikube("4096")

goFiles = local_output('find . -type f -name "*.go"').split('\n')

local_resource(
  'service-template-compile',
  'CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service-template ./',
  deps=goFiles,
  ignore=['**/*_test.go'],
)

legacy_directory = os.path.realpath('../legacy')

local_resource(
  'docker-compose-up-legacy',
  'PROJECT_DIR={} docker-compose -f {}/modanisa-config-files/docker-phalcon/docker-compose/docker-compose.yml up -d --build'.format(legacy_directory, legacy_directory)
)

docker_build_with_restart('service-template-image', '.',
  entrypoint='/app/service-template',
  dockerfile='Dockerfile',
  only=[
    './.config',
    'service-template',
  ],
  live_update=[
    sync('./service-template', '/app/service-template'),
  ])

k8s_yaml(['.dev/deployment.yaml'])

k8s_resource('service-template', port_forwards=3001)

local_resource(
  'picking_logs',
  serve_cmd = 'docker exec -i docker-compose_picking_1 /usr/bin/tail -f /tmp/jdip.log',
  resource_deps = ['docker-compose-up-legacy']
)

local_resource(
  'sorting_logs',
  serve_cmd = 'docker exec -i docker-compose_sorting_1 /usr/bin/tail -f /tmp/jdip.log',
  resource_deps = ['docker-compose-up-legacy']
)

local_resource(
  'accept-depo_logs',
  serve_cmd = 'docker exec -i docker-compose_accept-depo_1 /usr/bin/tail -f /tmp/jdip.log',
  resource_deps = ['docker-compose-up-legacy']
)

local_resource(
  'depo-yonetim_logs',
  serve_cmd = 'docker exec -i docker-compose_depo-yonetim_1 /usr/bin/tail -f /tmp/jdip.log',
  resource_deps = ['docker-compose-up-legacy']
)

local_resource(
  'depo-api_logs',
  serve_cmd = 'docker exec -i docker-compose_depo-api_1 /usr/bin/tail -f /tmp/jdip.log',
  resource_deps = ['docker-compose-up-legacy']
)