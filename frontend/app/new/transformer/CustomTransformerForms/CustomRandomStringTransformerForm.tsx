'use client';
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
} from '@/components/ui/form';

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

import { Switch } from '@/components/ui/switch';
import { ReactElement } from 'react';
import { useFormContext } from 'react-hook-form';

interface Props {
  isDisabled?: boolean;
}

export default function CustomRandomStringTransformerForm(
  props: Props
): ReactElement {
  const fc = useFormContext();

  const { isDisabled } = props;

  const digitLength = Array.from({ length: 12 }, (_, index) => index + 1);

  return (
    <div className="flex flex-col w-full space-y-4 pt-4">
      <FormField
        name={`config.config.value.preserveLength`}
        control={fc.control}
        render={({ field }) => (
          <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
            <div className="space-y-0.5">
              <FormLabel>Preserve Length</FormLabel>
              <FormDescription>
                Set the length of the output string to be the same as the input
              </FormDescription>
            </div>
            <FormControl>
              <Switch
                checked={field.value}
                onCheckedChange={field.onChange}
                disabled={
                  fc.watch('config.config.value.strLength') || isDisabled
                }
              />
            </FormControl>
          </FormItem>
        )}
      />
      <FormField
        name={`config.config.value.strLength`}
        control={fc.control}
        render={({ field }) => (
          <FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
            <div className="space-y-0.5">
              <FormLabel>String Length</FormLabel>
              <FormDescription>
                Set the length of the output string. The default length is 4.
                The max length is 18.
              </FormDescription>
            </div>
            <FormControl>
              <Select onValueChange={field.onChange} value={field.value}>
                <SelectTrigger
                  className="w-[180px]"
                  disabled={
                    fc.watch('config.config.value.preserveLength') || isDisabled
                  }
                >
                  <SelectValue placeholder="4" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    {digitLength.map((item) => (
                      <SelectItem value={String(item)} key={item}>
                        {item}
                      </SelectItem>
                    ))}
                  </SelectGroup>
                </SelectContent>
              </Select>
            </FormControl>
          </FormItem>
        )}
      />
    </div>
  );
}
