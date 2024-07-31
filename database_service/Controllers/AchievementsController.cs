using database_service.Database;
using database_service.Database.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace database_service.Controllers;

[Route("api/[controller]")]
[ApiController]
public class AchievementsController(DatabaseContext context) : ControllerBase
{
	private readonly DatabaseContext _context = context;

	// GET: api/<AchievementsController>
	[HttpGet]
	public IEnumerable<Achievement> Get()
	{
		return _context.Achievements;
	}

	// GET api/<AchievementsController>/5
	[HttpGet("{id}")]
	public async Task<Achievement?> Get(int id)
	{
		return await _context.Achievements.FirstOrDefaultAsync(a => a.Id == id);
	}

	// POST api/<AchievementsController>
	[HttpPost]
	public async Task<bool> Post([FromForm] Achievement achievement)
	{
		if (await _context.Achievements.AnyAsync(a => a.Id == achievement.Id))
			return false;

		await _context.Achievements.AddAsync(achievement);
		await _context.SaveChangesAsync();
		return true;
	}

	// PUT api/<AchievementsController>/5
	[HttpPut("{id}")]
	public async Task<bool> Put(int id, [FromForm] Achievement newAchievement)
	{
		var oldAchievement = await _context.Achievements.FirstOrDefaultAsync(a => a.Id == id);
		if (oldAchievement is null)
			return false;

		oldAchievement.Name = newAchievement.Name;
		oldAchievement.Description = newAchievement.Description;
		oldAchievement.Reward = newAchievement.Reward;

		await _context.SaveChangesAsync();
		return true;
    }

	// DELETE api/<AchievementsController>/5
	[HttpDelete("{id}")]
	public async Task<bool> Delete(int id)
	{
		var achievement = await _context.Achievements.FirstOrDefaultAsync(a => a.Id == id);
		if (achievement is null)
			return false;

		_context.Achievements.Remove(achievement);
		await _context.SaveChangesAsync();
		return true;
    }
}
