const Discord = require('discord.js');

const dotenv = require('dotenv');
dotenv.config();

// create the client and login
const client = new Discord.Client();
client.login(process.env.DISCORD);

// ready event
client.on('ready', () => {
    client.user.setActivity('with your emotions.');

    return console.log(`Bot is ready and has started on ${client.guilds.size} servers.`);
});

// message event
client.on('message', msg => {
    // if the author is a bot, ignore the message
    if (msg.author.bot) return;

    const commandPrefix = 'js!';

    // make a ping command
    if (msg.content === `${commandPrefix}ping`) {
        return msg.channel.send('pong!');
    }
});
