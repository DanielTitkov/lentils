package job

func (j *Job) GatherAndSendServiceStats() {
	// stats, err := j.app.ServiceStats()
	// if err != nil {
	// 	j.logger.Error("failed to gather service stats", err)
	// 	return
	// }

	// jsonStats, err := json.MarshalIndent(stats, "", "  ")
	// if err != nil {
	// 	j.logger.Error("failed to marshal service stats", err)
	// 	return
	// }

	// bot, err := tgbotapi.NewBotAPI(j.cfg.External.Telegram.TelegramToken)
	// if err != nil {
	// 	j.logger.Error("failed to create telegram connection", err)
	// 	return
	// }
	// defer bot.StopReceivingUpdates()

	// msg := tgbotapi.NewMessage(j.cfg.External.Telegram.TelegramTo, fmt.Sprintf(
	// 	"Service stats for %s:\n\n%s\n\nSent from '%s' enviroment",
	// 	time.Now().Format(time.RFC1123Z), jsonStats, j.cfg.Env,
	// ))

	// _, err = bot.Send(msg)
	// if err != nil {
	// 	j.logger.Error("failed to send message", err)
	// 	return
	// }

	// j.logger.Info("sent service stats", fmt.Sprintf("chat id: %d", j.cfg.External.Telegram.TelegramTo))
}
