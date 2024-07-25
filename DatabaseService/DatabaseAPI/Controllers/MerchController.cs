using DatabaseAPI.Database;
using DatabaseAPI.Database.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace DatabaseAPI.Controllers;

[Route("api/[controller]")]
[ApiController]
public class MerchController(DatabaseContext context) : ControllerBase
{
    private readonly DatabaseContext _context = context;

    // GET: api/<MerchController>
    [HttpGet]
    public IEnumerable<Merch> Get()
    {
        return _context.Merches;
    }

    // GET api/<MerchController>/5
    [HttpGet("{id}")]
    public async Task<Merch?> Get(int id)
    {
        return await _context.Merches.FirstOrDefaultAsync(m => m.Id == id);
    }

    // POST api/<MerchController>
    [HttpPost]
    public async Task<bool> Post([FromForm] Merch merch)
    {
        if (await _context.Merches.AnyAsync(m => m.Id == merch.Id))
            return false;

        await _context.Merches.AddAsync(merch);
        await _context.SaveChangesAsync();
        return true;
    }

    // PUT api/<MerchController>/5
    [HttpPut("{id}")]
    public async Task<bool> Put(int id, [FromBody] Merch newMerch)
    {
        var oldMerch = await _context.Merches.FirstOrDefaultAsync(m => m.Id == id);
        if (oldMerch is null)
            return false;

        oldMerch.Name = newMerch.Name;
        oldMerch.Description = newMerch.Description;
        oldMerch.Price = newMerch.Price;

        await _context.SaveChangesAsync();
        return true;
    }

    // DELETE api/<MerchController>/5
    [HttpDelete("{id}")]
    public async Task<bool> Delete(int id)
    {
        var merch = await _context.Merches.FirstOrDefaultAsync(m => m.Id == id);
        if (merch is null)
            return false;

        _context.Merches.Remove(merch);
        await _context.SaveChangesAsync();
        return true;
    }
}
