# DiscordHA (Discord High Availability)

DiscordHA is library on to be used together with [discordgo](https://github.com/bwmarrin/discordgo) to deploy Discord bots in high availability.
This will allow deploying Discord bots with zero-downtime and without the risk of handling events twice. DiscordHA is designed to run in a Kunernetes environment with multiple replica's available. 
It relies on Etcd as a locking system to prevent events of being received twice, this works in a first locked principle for timestamped events (e.g. messages), and a leader election system for non timestamped events (e.g. reactions).

DiscordHA is not meant for sharding but enables Discord bots to have multiple replicas wtahcing the same guilds.

## Example

```go
func main(){
	dg, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		fmt.Prinf("error creating Discord session: %v\n", err)
        return
	}

	haLogger := log.New(os.Stderr, "discordha: ", log.Ldate|log.Ltime)
	ha, err := discordha.New(&discordha.Config{
		Session:       dg,
		HA:            true, // set to false to run 1 replica for debugging, this disables locking and caching
		EtcdEndpoints: []string{"etcd-bob.etcd.svc.cluster.local:2379"},
		Context:       context.TODO(),
		Log:           *haLogger,
	})
	if err != nil {
        fmt.Prinf("error creating Discord HA: %v\n", err)
		return
	}
    defer s.ha.Stop()
    
    // using AddHandler on discordha will handle it on one replica
	ha.AddHandler(func(sess *discordgo.Session, m *discordgo.MessageCreate) {
        fmt.Printf("Received Message: %q", m.Message.Content)
    })
   

	err = dg.Open()
	if err != nil {
        fmt.Prinf("error opening connectio: %v\n", err)
        return
	}
    defer dg.Close()

	log.Println("DiscordHA Example is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
```

## Used by
- [Thomas Bot](https://github.com/itfactory-tm/thomas-bot)

## Roadmap
- [ ] Add support for sharded event listeners
- [ ] Improve the voice system
- [ ] Improve locking system for events
- [ ] Facilitate Etcd over HTTPS