using DatabaseAPI.Database;
using DatabaseAPI.Database.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System.Text.RegularExpressions;

namespace DatabaseAPI.Controllers;

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

	// POST api/<CoinsRequestsController>
	[HttpPost]
	public async Task<bool> Post([FromForm] CoinsRequest coinsRequest)
	{
        if (await _context.CoinsRequests.AnyAsync(cr => cr.Id == coinsRequest.Id))
            return false;

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
