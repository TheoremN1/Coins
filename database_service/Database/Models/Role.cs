using System.ComponentModel.DataAnnotations;

namespace database_service.Database.Models;

public class Role
{
	[Key]
	public string Key { get; set; }

	public string Name { get; set; }
}
