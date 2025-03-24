#!/usr/bin/env python3
import asyncio
import discord
import logging
import sys
from os import getenv

from dotenv import load_dotenv
from discord.ext import commands

file_handler = logging.FileHandler(filename="bot.log", encoding="utf-8", mode="w")
stream_handler = logging.StreamHandler(sys.stdout)
logger = logging.getLogger("discord")
logger.setLevel(logging.INFO)
logger.addHandler(file_handler)
logger.addHandler(stream_handler)

client = commands.Bot(
    command_prefix="!",
    intents=discord.Intents(voice_states=True, messages=True, guilds=True),
)


@client.event
async def on_guild_join(guild: discord.Guild):
    logger.info(f"Joined guild: {guild.name}")


load_dotenv()


async def init():
    logger.info("Loading extensions...")
    await client.load_extension("joinleave")
    logger.info("Loaded extension.")


if __name__ == "__main__":
    asyncio.run(init())

    logger.info("Starting bot...")
    client.run(getenv("BOT_TOKEN"))
