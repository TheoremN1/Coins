using DatabaseAPI.Database;
using DatabaseAPI.Database.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace DatabaseAPI.Controllers;

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
