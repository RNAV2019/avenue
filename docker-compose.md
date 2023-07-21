Start the docker containers:
docker-compose -f ./supabase/docker/docker-compose.yml -f ./docker-compose.yml up

Stop the docker containers:
docker-compose -f ./supabase/docker/docker-compose.yml -f ./docker-compose.yml down -v --remove-orphans