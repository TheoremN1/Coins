using System.ComponentModel.DataAnnotations;

namespace DatabaseAPI.Database.Models;

public class Role
{
	[Key]
	public string Key { get; set; }

	public string Name { get; set; }
}
