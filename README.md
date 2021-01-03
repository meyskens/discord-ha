# DiscordHA (Discord High Availability)

DiscordHA is library on to be used together with [discordgo](https://github.com/bwmarrin/discordgo) to deploy Discord bots in high availability.
This will allow deploying Discord bots with zero-downtime and without the risk of handling events twice. DiscordHA is designed to run in a Kunernetes environment with multiple replica's available. 
It relies on Etcd as a locking system to prevent events of being received twice, this works in a first locked principle for timestamped events (e.g. messages), and a leader election system for non timestamped events (e.g. reactions).

DiscordHA is not meant for sharding but enables Discord bots to have multiple replicas wtahcing the same guilds.

## Used by
- [Thomas Bot](https://github.com/itfactory-tm/thomas-bot)

## Roadmap
- [ ] Add support for sharded event listeners
- [ ] Improve the voice system
- [ ] Improve locking system for events
