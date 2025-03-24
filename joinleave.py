from discord.ext import commands


class JoinLeave(commands.Cog):
    def __init__(self, bot):
        self.bot = bot

    @commands.Cog.listener()
    async def on_voice_state_update(self, member, before, after):
        jlto = 120
        member_name = f"{member.display_name} ({member})"
        if before.channel is None and after.channel is not None:
            await after.channel.send(
                content=f"{member_name} joined.", delete_after=jlto
            )
        elif before.channel is not None and after.channel is None:
            await before.channel.send(content=f"{member_name} left.", delete_after=jlto)
        elif (
            before.channel is not None
            and after.channel is not None
            and before.channel.id != after.channel.id
        ):
            await before.channel.send(content=f"{member_name} left.", delete_after=jlto)
            await after.channel.send(
                content=f"{member_name} joined.", delete_after=jlto
            )


async def setup(bot):
    await bot.add_cog(JoinLeave(bot))
