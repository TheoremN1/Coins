using DatabaseAPI.Database;
using DatabaseAPI.Database.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace DatabaseAPI.Controllers;

[Route("api/[controller]")]
[ApiController]
public class RolesController(DatabaseContext context) : ControllerBase
{
	private readonly DatabaseContext _context = context;

	// GET: api/<RolesController>
	[HttpGet]
	public IEnumerable<Role> Get()
	{
		return _context.Roles;
	}

	// GET api/<RolesController>/user
	[HttpGet("{key}")]
	public async Task<Role?> Get(string key)
	{
		return await _context.Roles.FirstOrDefaultAsync(r => r.Key == key);
	}
}
