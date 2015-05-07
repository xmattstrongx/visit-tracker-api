PROJECT_ROOT = Dir.pwd
DOCKER_DIR = "#{PROJECT_ROOT}/docker"
SERVICES_DIR = "#{DOCKER_DIR}/services"
SERVICES = `ls #{SERVICES_DIR}`.split()
OUTPUT_DIR = "#{DOCKER_DIR}/bin"
BUILD_FILE = 'Rakefile'
IMAGE_PREFIX = 'mstrong'
REGISTRY_URL = 'mytest.docker.registry'
OUTPUTS = SERVICES.map{|name| File.join(OUTPUT_DIR, name)}
CLEAN.include(OUTPUT_DIR)

task :default => [:build]

desc "Build all services"
task :build => SERVICES.each {|service| service}

desc "Build all services concurrently"
multitask :multibuild => SERVICES.each {|service| service}

task :check_docker do
  fail_if_cmd_fails('which docker', "You don't seem to have docker installed. Stopping...")
end

task :check_conn do
  fail_if_cmd_fails('docker ps', "Something appears to be wrong with your Docker connection. Stopping...")
end

desc "Start the project"
task :start => [:check_docker, :check_conn] do
  sh "docker-compose up -d"
end

task :push => [:check_docker, :check_conn] do
  SERVICES.each do |service|
    image_name = "#{IMAGE_PREFIX}-#{service}"
    sh "docker tag -f #{image_name} #{REGISTRY_URL}/#{image_name}"
    sh "docker push #{REGISTRY_URL}/#{image_name}"
  end
end

def build_service_image(service_name, service_dir, output_dir)
  # copy all of the files except the builder file
  files = `ls #{service_dir}`
    .split().select{|file| file != BUILD_FILE}
    .each{|file| sh "cp #{service_dir}/#{file} #{output_dir}/#{file}"}

  # build the docker image
  sh "docker build -t #{IMAGE_PREFIX}-#{service_name} #{output_dir}"
end

# Create a build task for each service
SERVICES.each do |service|
  desc "Build #{service}"
  task service => [:check_docker, :check_conn] do
    # build any artifacts for this service
    # run the build file for this service if it exists
    service_dir = "#{SERVICES_DIR}/#{service}"
    build_filename = "#{service_dir}/#{BUILD_FILE}"
    output_dir = "#{OUTPUT_DIR}/#{service}"
    sh "mkdir -p #{output_dir}"
    if File.exist?(build_filename)
      sh "rake -f #{service_dir}/#{BUILD_FILE} default[#{PROJECT_ROOT},#{output_dir}]"
    end
    # now build the image for this service
    build_service_image(service, service_dir, output_dir)
  end
end

def fail_if_cmd_fails(cmd, msg)
  output = `#{cmd}`
  if $?.exitstatus != 0
    puts output
    fail msg
  end
  return output
end
