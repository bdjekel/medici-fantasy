#!/bin/bash
echo "Updating .env.example..."
sed 's/=.*/=<CHANGE_ME>/' .env > .env.example
echo ".env.example updated!"
