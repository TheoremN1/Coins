using DatabaseAPI.Database;
using DatabaseAPI.Database.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace DatabaseAPI.Controllers;

[Route("api/[controller]")]
[ApiController]
public class UsersController(DatabaseContext context, ILogger<UsersController> logger) : ControllerBase
{
	private readonly DatabaseContext _context = context;
	private readonly ILogger<UsersController> _logger = logger;

	// GET: api/<UsersController>
	[HttpGet]
	public IEnumerable<User> Get()
	{
        _logger.LogInformation($"GET /api/users");
        return _context.Users;
	}

	// GET api/<UsersController>/5
	[HttpGet("{id}")]
	public async Task<User?> Get(int id)
	{
        _logger.LogInformation($"GET /api/users/{id}");
        return await _context.Users.FirstOrDefaultAsync(u => u.Id == id);
	}

    // GET api/<UsersController>/5/role
    [HttpGet("{id}/role")]
    public async Task<Role?> GetRole(int id)
    {
        _logger.LogInformation($"GET /api/users/{id}/role");
        var user = await Get(id);
        return user?.Role;
    }

    // POST api/<UsersController>
    [HttpPost]
	public async Task<bool> Post([FromForm] User user)
	{
		_logger.LogInformation($"POST /api/users Data: {user}");

		if (await _context.Users.AnyAsync(u => u.Login == user.Login || u.Id == user.Id))
			return false;

		await _context.Users.AddAsync(user);
		await _context.SaveChangesAsync();
        return true;
	}

	// PUT api/<UsersController>/5
	[HttpPut("{id}")]
	public async Task<bool> Put(int id, [FromForm] User newUser)
	{
        var oldUser = await _context.Users.FirstOrDefaultAsync(u => u.Id == id);
		if (oldUser is null)
			return false;

        _logger.LogInformation($"PUT /api/users \nOldData: {newUser}\nNewData: {newUser}");

        if (await _context.Users.AnyAsync(u => u.Login == newUser.Login && u.Id != id))
			return false;

		oldUser.Name = newUser.Name;
		oldUser.Surname = newUser.Surname;
		oldUser.Balance = newUser.Balance;
		oldUser.Login = newUser.Login;
		oldUser.Password = newUser.Password;
		oldUser.RoleKey = newUser.RoleKey;

		await _context.SaveChangesAsync();
		return true;
	}

	// DELETE api/<UsersController>/5
	[HttpDelete("{id}")]
	public async Task<bool> Delete(int id)
	{
        _logger.LogInformation($"DELETE /api/users/{id}");

        var user = await _context.Users.FirstOrDefaultAsync(u => u.Id == id);
		if (user is null)
			return false;

		_context.Users.Remove(user);
		await _context.SaveChangesAsync();
		return true;
	}
}
