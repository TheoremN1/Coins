using Microsoft.AspNetCore.Mvc;
using DatabaseAPI.Database;
using DatabaseAPI.Database.Models;

namespace DatabaseAPI.Controllers;

[Route("api/[controller]")]
[ApiController]
public class UsersController(DatabaseContext context) : ControllerBase
{
    private readonly DatabaseContext _context = context;

    // GET: api/<UsersController>
    [HttpGet]
    public IEnumerable<User> Get()
    {
        return _context.Users;
    }

    // GET api/<UsersController>/5
    [HttpGet("{id}")]
    public User? Get(int id)
    {
        return _context.Users.FirstOrDefault(u => u.Id == id);
    }

    // POST api/<UsersController>
    [HttpPost]
    public bool Post([FromForm] User user)
    {
        if (_context.Users.Any(u => u.Login == user.Login || u.Id == user.Id))
            return false;

        _context.Users.Add(user);
        _context.SaveChanges();
        return true;
    }

    // PUT api/<UsersController>/5
    [HttpPut("{id}")]
    public bool Put(int id, [FromForm] User newUser)
    {
        var oldUser = _context.Users.FirstOrDefault(u => u.Id == id);
        if (oldUser is null)
            return false;

        if (_context.Users.Any(u => u.Login == newUser.Login && u.Id != id))
            return false;

        oldUser.Name = newUser.Name;
        oldUser.Surname = newUser.Surname;
        oldUser.Balance = newUser.Balance;
        oldUser.Login = newUser.Login;
        oldUser.Password = newUser.Password;
        oldUser.RoleKey = newUser.RoleKey;

        _context.SaveChanges();
        return true;
    }

    // DELETE api/<UsersController>/5
    [HttpDelete("{id}")]
    public bool Delete(int id)
    {
        var user = _context.Users.FirstOrDefault(u => u.Id == id);
        if (user is null)
            return false;

        _context.Users.Remove(user);
        _context.SaveChanges();
        return true;
    }
}
