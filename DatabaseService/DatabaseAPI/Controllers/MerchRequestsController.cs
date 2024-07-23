using Microsoft.AspNetCore.Mvc;

namespace DatabaseAPI.Controllers;

[Route("api/[controller]")]
[ApiController]
public class MerchRequestsController : ControllerBase
{
    // GET: api/<MerchRequestsController>
    [HttpGet]
    public IEnumerable<string> Get()
    {
        return new string[] { "value1", "value2" };
    }

    // GET api/<MerchRequestsController>/5
    [HttpGet("{id}")]
    public string Get(int id)
    {
        return "value";
    }

    // POST api/<MerchRequestsController>
    [HttpPost]
    public void Post([FromBody] string value)
    {
    }

    // PUT api/<MerchRequestsController>/5
    [HttpPut("{id}")]
    public void Put(int id, [FromBody] string value)
    {
    }

    // DELETE api/<MerchRequestsController>/5
    [HttpDelete("{id}")]
    public void Delete(int id)
    {
    }
}
