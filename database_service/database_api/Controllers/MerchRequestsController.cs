using DatabaseAPI.Database;
using DatabaseAPI.Database.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace DatabaseAPI.Controllers;

[Route("api/[controller]")]
[ApiController]
public class MerchRequestsController(DatabaseContext context) : ControllerBase
{
    private readonly DatabaseContext _context = context;

    // GET: api/<MerchRequestsController>
    [HttpGet]
    public IEnumerable<MerchRequest> Get()
    {
        return _context.MerchRequests;
    }

    // GET api/<MerchRequestsController>/5
    [HttpGet("{id}")]
    public async Task<MerchRequest?> Get(int id)
    {
        return await _context.MerchRequests.FirstOrDefaultAsync(mr => mr.Id == id);
    }

    // GET api/<MerchRequestsController>/5/user
    [HttpGet("{id}/user")]
    public async Task<User?> GetUser(int id)
    {
        var merchRequest = await Get(id);
        return merchRequest?.User;
    }

    // GET api/<MerchRequestsController>/5/hr
    [HttpGet("{id}/hr")]
    public async Task<User?> GetHr(int id)
    {
        var merchRequest = await Get(id);
        return merchRequest?.Hr;
    }

    // GET api/<MerchRequestsController>/5/merch
    [HttpGet("{id}/merch")]
    public async Task<Merch?> GetMerch(int id)
    {
        var merchRequest = await Get(id);
        return merchRequest?.Merch;
    }

    // GET api/<MerchRequestsController>/5/status
    [HttpGet("{id}/status")]
    public async Task<Status?> GetStatus(int id)
    {
        var merchRequest = await Get(id);
        return merchRequest?.Status;
    }

    // POST api/<MerchRequestsController>
    [HttpPost]
    public async Task<bool> Post([FromForm] MerchRequest merchRequest)
    {
        if (await _context.MerchRequests.AnyAsync(mr => mr.Id == merchRequest.Id))
            return false;

        if(merchRequest.HrId < 1)
        {
            merchRequest.HrId = null;
            merchRequest.HrMessage = null;
        }

        await _context.MerchRequests.AddAsync(merchRequest);
        await _context.SaveChangesAsync();
        return true;
    }

    // PUT api/<MerchRequestsController>/5
    [HttpPut("{id}")]
    public async Task<bool> Put(int id, [FromForm] MerchRequest newMerchRequest)
    {
        var oldMerchRequest = await _context.MerchRequests.FirstOrDefaultAsync(cr => cr.Id == id);
        if (oldMerchRequest is null)
            return false;

        if (newMerchRequest.HrId < 1)
        {
            newMerchRequest.HrId = null;
            newMerchRequest.HrMessage = null;
        }

        oldMerchRequest.UserMessage = newMerchRequest.UserMessage;
        oldMerchRequest.HrId = newMerchRequest.HrId;
        oldMerchRequest.HrMessage = newMerchRequest.HrMessage;
        oldMerchRequest.MerchId = newMerchRequest.MerchId;
        oldMerchRequest.StatusKey = newMerchRequest.StatusKey;

        await _context.SaveChangesAsync();
        return true;
    }

    // DELETE api/<MerchRequestsController>/5
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
