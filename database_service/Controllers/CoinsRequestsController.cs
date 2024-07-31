using database_service.Database;
using database_service.Database.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System.Text.RegularExpressions;

namespace database_service.Controllers;

[Route("api/[controller]")]
[ApiController]
public class CoinsRequestsController(DatabaseContext context) : ControllerBase
{
    private readonly DatabaseContext _context = context;

    // GET: api/<CoinsRequestsController>
    [HttpGet]
	public IEnumerable<CoinsRequest> Get()
	{
		return _context.CoinsRequests;
	}

	// GET api/<CoinsRequestsController>/5
	[HttpGet("{id}")]
	public async Task<CoinsRequest?> Get(int id)
	{
		return await _context.CoinsRequests.FirstOrDefaultAsync(cr => cr.Id == id);
	}

    // GET api/<CoinsRequestsController>/5/user
    [HttpGet("{id}/user")]
    public async Task<User?> GetUser(int id)
    {
        var coinsRequest = await Get(id);
        return coinsRequest?.User;
    }

    // GET api/<CoinsRequestsController>/5/hr
    [HttpGet("{id}/hr")]
    public async Task<User?> GetHr(int id)
    {
        var coinsRequest = await Get(id);
        return coinsRequest?.Hr;
    }

    // GET api/<CoinsRequestsController>/5/achievement
    [HttpGet("{id}/achievement")]
    public async Task<Achievement?> GetMerch(int id)
    {
        var coinsRequest = await Get(id);
        return coinsRequest?.Achievement;
    }

    // GET api/<CoinsRequestsController>/5/status
    [HttpGet("{id}/status")]
    public async Task<Status?> GetStatus(int id)
    {
        var coinsRequest = await Get(id);
        return coinsRequest?.Status;
    }

    // POST api/<CoinsRequestsController>
    [HttpPost]
	public async Task<bool> Post([FromForm] CoinsRequest coinsRequest)
	{
        if (await _context.CoinsRequests.AnyAsync(cr => cr.Id == coinsRequest.Id))
            return false;

        if (coinsRequest.HrId < 1)
        {
            coinsRequest.HrId = null;
            coinsRequest.HrMessage = null;
        }

        await _context.CoinsRequests.AddAsync(coinsRequest);
        await _context.SaveChangesAsync();
        return true;
    }

	// PUT api/<CoinsRequestsController>/5
	[HttpPut("{id}")]
	public async Task<bool> Put(int id, [FromForm] CoinsRequest newCoinsRequest)
	{
        var oldCoinsRequest = await _context.CoinsRequests.FirstOrDefaultAsync(cr => cr.Id == id);
        if (oldCoinsRequest is null)
            return false;

        if (newCoinsRequest.HrId < 1)
        {
            newCoinsRequest.HrId = null;
            newCoinsRequest.HrMessage = null;
        }

        oldCoinsRequest.UserMessage = newCoinsRequest.UserMessage;
        oldCoinsRequest.HrId = newCoinsRequest.HrId;
        oldCoinsRequest.HrMessage = newCoinsRequest.HrMessage;
        oldCoinsRequest.AchievementId = newCoinsRequest.AchievementId;
        oldCoinsRequest.StatusKey = newCoinsRequest.StatusKey;

        await _context.SaveChangesAsync();
        return true;
    }

	// DELETE api/<CoinsRequestsController>/5
	[HttpDelete("{id}")]
	public async Task<bool> Delete(int id)
	{
        var coinsRequest = await _context.CoinsRequests.FirstOrDefaultAsync(cr => cr.Id == id);
        if (coinsRequest is null)
            return false;

        _context.CoinsRequests.Remove(coinsRequest);
        await _context.SaveChangesAsync();
        return true;
    }
}
