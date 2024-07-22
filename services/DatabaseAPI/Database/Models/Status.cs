using System.ComponentModel.DataAnnotations;

namespace DatabaseAPI.Database.Models;

public class Status
{
	[Key]
	public string Key { get; set; }

	public string Name { get; set; }
}