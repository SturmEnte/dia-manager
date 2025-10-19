# dia-manager

## How to use

I intented the project to be used with Docker. If you don't plan on using Docker this guide is not for you.
You already need to have a postgres database running.

1. Execute the SQL in `DB.SQL` in the database you want to use for the project
2. Clone the project on your server
   `git clone https://github.com/SturmEnte/dia-manager.git`
3. cd into the project
   `cd dia-manager`
4. Create the .env file by copying it and then setting the values to the ones needed for your setup
   `cp .env.example .env`
   `nano .env`
5. Build and deploy the project
   `docker compose up -d --build`
