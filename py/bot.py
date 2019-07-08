# import discord.py
import discord
from discord.ext import commands

# load dotenv
from dotenv import load_dotenv
load_dotenv()

# other imports
import os

# create the bot
bot = commands.Bot(command_prefix='py!')

# ready event
@bot.event
async def on_ready():
    await bot.change_presence(status=discord.Status.online, activity=discord.Game("with your emotions"))
    print('Bot is ready and has started on ' + str(len(bot.guilds)) + ' servers.')

# ping command
@bot.command()
async def ping(ctx):
    await ctx.send('pong!')

# login
bot.run(os.getenv('DISCORD'))