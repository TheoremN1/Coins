using database_service.Database;
using database_service.Database.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace database_service.Controllers;

[Route("api/[controller]")]
[ApiController]
public class StatusesController(DatabaseContext context) : ControllerBase
{
	private readonly DatabaseContext _context = context;

	// GET: api/<StatusesController>
	[HttpGet]
	public IEnumerable<Status> Get()
	{
		return _context.Statuses;
	}

	// GET api/<StatusesController>/wait
	[HttpGet("{key}")]
	public async Task<Status?> Get(string key)
	{
		return await _context.Statuses.FirstOrDefaultAsync(s => s.Key == key);
	}
}
