using System.ComponentModel.DataAnnotations;

namespace database_service.Database.Models;

public class Status
{
	[Key]
	public string Key { get; set; }

	public string Name { get; set; }
}