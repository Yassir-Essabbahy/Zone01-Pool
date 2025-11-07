curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq  ".[] | select(.id== $HERO_ID)" |jq ".connections.relatives" | tr -d '"'

